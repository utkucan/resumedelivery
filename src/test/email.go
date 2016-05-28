package hello

import (
    "fmt"
  	"net/http"

  	// "appengine"
   //  "appengine/log"
  	// "appengine/mail"
)

func init() {
  	http.HandleFunc("/send", sendEmail)
}


func sendEmail(w http.ResponseWriter, r *http.Request) {
				//  fmt.Println(“got:”, r.URL.Query());
        fmt.Printf("hello, world\n")
  			// toEmail := r.URL.Query()["to"][0];
    		fmt.Fprint(w, "asdasddas")
  
  			// ctx := appengine.NewContext(r)
     //    url := "asdasdasd"
     //    msg := &mail.Message{
     //            Sender:  "<utkucanyucel@gmail.com>",
     //            To:      []string{toEmail},
     //            Subject: "Bu Bir OTOMATIK E-postadır",
     //            Body:    fmt.Sprintf("test", url),
     //    }
  			fmt.Fprint(w, "geldi")
      //   if err := mail.Send(ctx, msg); err != nil {
						// log.Errorf(ctx, "Couldn't send email: %v", err)
      //   }else{
      //     	fmt.Fprint(w, "gonderildi.")
      //   }
}

//test