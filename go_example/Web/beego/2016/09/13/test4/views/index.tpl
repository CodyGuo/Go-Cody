<!DOCTYPE html>
<html>
  <head>
    <title>VIP 视频播放器</title>
    <meta name="viewport" content="width=device-width,initial-scale=1.0,
        minimum-scale=1.0,maximum-scale=1.0,user-scalable=no" >
      <meta name="apple-itunes-app" content="" />
      <meta name="format-detection" content="telephone=no, address=no" >
      <meta name="apple-mobile-web-app-capable" content="yes" />
      <meta name="apple-mobile-web-app-status-bar-style" content="black-translucent" />
      <meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
      <meta http-equiv="pragma" content="no-cache">
    <meta http-equiv="cache-control" content="no-cache,must-revalidate">
    <meta http-equiv="expires" content="0">
      <meta http-equiv="keywords" content="Joyce Player">
      <meta http-equiv="description" content="Joyce Player">
      <meta http-equiv="content-type" content="application/xhtml+xml;charset=UTF-8">

    <style type="text/css">
      .container {
        width: 1024px;
        height: auto;
        margin: 0px auto;
      }
      .action {
        width: 100%;
        height: 40px;
        text-align: center;
      }
      .player {
        width: 100%;
        height: auto;
      }
    </style>
  </head>

  <body>
    <div class="container">
      <div class="action">
        请输入视频地址: <input type="url" value="http://www.mgtv.com/v/2/104817/f/3426880.html" class="videoUrl" size="90%" autofocus="autofocus" placeholder="请输入视频地址">
        <input type="button" style="margin-left: 10px;" value="播放视频" onclick="javascript:play();">
      </div>
      <div class="player">
        <iframe class="play_container" width="100%" height="600"
          allowTransparency="true" frameborder="0" scrolling="no"></iframe>
      </div>
    </div>
    <script src="http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js"></script>
    <script type="text/javascript">
      !window.jQuery && document.write('<script src="jquery-1.8.3.min.js"><\/script>');

      /*
       * qyz 21 -> http://www.mgtv.com/v/2/104817/f/3405841.html
       * qyz 22 -> http://www.mgtv.com/v/2/104817/f/3426880.html
       */
      var play = (function(){
        var $video_url = $('.videoUrl'),
          $play_container = $('.play_container');
        console.log($video_url);
        console.log($play_container);
        $play_container.attr('src', 'http://jxapi.nepian.com/ckparse/?url=' + $video_url.val());
      });
    </script>
  </body>
</html>