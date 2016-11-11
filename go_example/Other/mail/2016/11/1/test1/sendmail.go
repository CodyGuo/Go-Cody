func SendMail(addr string, a Auth, from string, to []string, msg []byte) error {
   c, err := Dial(addr)
   if err != nil {
      return err
   }
   defer c.Close()
   if err = c.hello(); err != nil {
      return err
   }
   if ok, _ := c.Extension("STARTTLS"); ok {
      config := &tls.Config{ServerName: c.serverName}
      if testHookStartTLS != nil {
         testHookStartTLS(config)
      }
      if err = c.StartTLS(config); err != nil {
         return err
      }
   }
   if a != nil && c.ext != nil {
      if _, ok := c.ext["AUTH"]; ok {
         if err = c.Auth(a); err != nil {
            return err
         }
      }
   }
   if err = c.Mail(from); err != nil {
      return err
   }
   for _, addr := range to {
      if err = c.Rcpt(addr); err != nil {
         return err
      }
   }
   w, err := c.Data()
   if err != nil {
      return err
   }
   _, err = w.Write(msg)
   if err != nil {
      return err
   }
   err = w.Close()
   if err != nil {
      return err
   }
   return c.Quit()
  }