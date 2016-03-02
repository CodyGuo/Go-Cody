package declarative

import (
	"github.com/lxn/walk"
)

func (mw MainWindow) CreateCody() error {
	w, err := walk.NewMainWindowCody()
	if err != nil {
		return err
	}

	tlwi := topLevelWindowInfo{
		Name:             mw.Name,
		Font:             mw.Font,
		ToolTipText:      "",
		MinSize:          mw.MinSize,
		MaxSize:          mw.MaxSize,
		ContextMenuItems: mw.ContextMenuItems,
		OnKeyDown:        mw.OnKeyDown,
		OnKeyPress:       mw.OnKeyPress,
		OnKeyUp:          mw.OnKeyUp,
		OnMouseDown:      mw.OnMouseDown,
		OnMouseMove:      mw.OnMouseMove,
		OnMouseUp:        mw.OnMouseUp,
		OnSizeChanged:    mw.OnSizeChanged,
		DataBinder:       mw.DataBinder,
		Layout:           mw.Layout,
		Children:         mw.Children,
	}

	builder := NewBuilder(nil)

	w.SetSuspended(true)
	builder.Defer(func() error {
		w.SetSuspended(false)
		return nil
	})

	builder.deferBuildMenuActions(w.Menu(), mw.MenuItems)

	return builder.InitWidget(tlwi, w, func() error {
		if len(mw.ToolBar.Items) > 0 {
			var tb *walk.ToolBar
			if mw.ToolBar.AssignTo == nil {
				mw.ToolBar.AssignTo = &tb
			}

			if err := mw.ToolBar.Create(builder); err != nil {
				return err
			}

			old := w.ToolBar()
			w.SetToolBar(*mw.ToolBar.AssignTo)
			old.Dispose()
		} else {
			builder.deferBuildActions(w.ToolBar().Actions(), mw.ToolBarItems)
		}

		if err := w.SetTitle(mw.Title); err != nil {
			return err
		}

		if err := w.SetSize(mw.Size.toW()); err != nil {
			return err
		}
		// 设置屏幕居中
		w.SetScreenCenter(true)

		imageList, err := walk.NewImageList(walk.Size{16, 16}, 0)
		if err != nil {
			return err
		}
		w.ToolBar().SetImageList(imageList)

		if mw.OnDropFiles != nil {
			w.DropFiles().Attach(mw.OnDropFiles)
		}

		if mw.AssignTo != nil {
			*mw.AssignTo = w
		}

		builder.Defer(func() error {
			w.Show()

			return nil
		})

		return nil
	})
}

func (mw MainWindow) RunCody() (int, error) {
	var w *walk.MainWindow

	if mw.AssignTo == nil {
		mw.AssignTo = &w
	}

	if err := mw.CreateCody(); err != nil {
		return 0, err
	}
	(*mw.AssignTo).SetScreenCenter(true)

	return (*mw.AssignTo).Run(), nil
}
