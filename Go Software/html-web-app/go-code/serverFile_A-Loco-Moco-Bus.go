# A Loco Moco Bus ~

package main

// .
  
import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
    
    
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}



func app_welcome_center_page() {
    
    
}



// . appHandler
func appHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "~ Selfie_Lunch // - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
  
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  if pagePath == "/account" {
      pageTitle = "Account Page"
      pageList = pageList
  }
  
  if pagePath == "/profile" {
      pageTitle = "Profile Page"
      pageList = pageList
  }
  
  
  
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }
  
  if pagePath == "/settings" {
      pageTitle = "Settings Page"
      pageList = pageList
  }
 
  
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler



// . pageHandler, ~ for User Pages °
func pageHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +

  pageTitle := "~ A Loco Moco Bus - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  if pagePath == "/account" {
      pageTitle = "Account Page"
      pageList = pageList
  }
  
  if pagePath == "/profile" {
      pageTitle = "Profile Page"
      pageList = pageList
  }
  
  
// ,  ° . +
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }
  
  
   #### Stuff ~
 + _ Driver
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Driver" {
      pageTitle = "Driver Page"
      pageList = pageList
  }
  
  
 + _ Equipment 
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Equipment" {
      pageTitle = "Equipment Page"
      pageList = pageList
  }
  
  
 + _ Ingredients 
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Ingredients" {
      pageTitle = "Ingredients Page"
      pageList = pageList
  }
  
  
 + _ Vehicle 
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Vehicle" {
      pageTitle = "Vehicle Page"
      pageList = pageList
  }
  
  
 
 
 #### Stuff ~
 + _ Schedule 
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Schedule" {
      pageTitle = "Schedule Page"
      pageList = pageList
  }
  
  
 + _ Menu
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Menu" {
      pageTitle = "Menu Page"
      pageList = pageList
  }
  
  
 + _ Branding 
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Branding" {
      pageTitle = "Branding Page"
      pageList = pageList
  }
  
  
 + _ Customers
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Customers" {
      pageTitle = "Customers Page"
      pageList = pageList
  }
  
 
 #### Stuff ~
 + _ Asana
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Asana" {
      pageTitle = "Asana Page"
      pageList = pageList
  }
  
  
 + _ Slack
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Slack" {
      pageTitle = "Slack Page"
      pageList = pageList
  }
  
  
 + _ Discord 
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Discord" {
      pageTitle = "Discord Page"
      pageList = pageList
  }
  
  
 + _ GumRoad
    // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/GumRoad" {
      pageTitle = "GumRoad Page"
      pageList = pageList
  }
 
 
#### Stuff ~
+ _ GitHub
   // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/GitHub" {
      pageTitle = "GitHub Page"
      pageList = pageList
  }
  
+ _ Google Cloud
   // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Google_Cloud" {
      pageTitle = "Google_Cloud Page"
      pageList = pageList
  }
  
+ _ Linktree
   // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Linktree" {
      pageTitle = "Linktree Page"
      pageList = pageList
  }
  
  
+ _ Printful
   // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Printful" {
      pageTitle = "Printful Page"
      pageList = pageList
  }
  
  
    // ,  ° . + ` <p>Cafe Needs<ul>
    
    
   // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Menu" {
      pageTitle = "Menu Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Location ~
  // ,  ° . +
  if pagePath == "/Location" {
      pageTitle = "Location Page"
      pageList = pageList
  }
  
  
   // ,  ° . + - _ ` Equiptment ~
  // ,  ° . +
  if pagePath == "/Equiptment" {
      pageTitle = "Equiptment Page"
      pageList = pageList
  }
  
  
    // ,  ° . +- _ ` Staff ~
  // ,  ° . +
  if pagePath == "/Staff" {
      pageTitle = "Staff Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Services ~
  // ,  ° . +
  if pagePath == "/Services" {
      pageTitle = "Services Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Music ~
  // ,  ° . +
  if pagePath == "/Music" {
      pageTitle = "Music Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Setting ~
  // ,  ° . +
  if pagePath == "/Setting" {
      pageTitle = "Setting Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Sources ~
  // ,  ° . +
  if pagePath == "/Sources" {
      pageTitle = "Sources Page"
      pageList = pageList
  }
  
  
   // ,  ° . +  - _ ` Marketing ~
  // ,  ° . +
  if pagePath == "/Marketing" {
      pageTitle = "Marketing Page"
      pageList = pageList
  }
  
  
  // ,  ° . +  - _ ` Operations ~
  // ,  ° . +
  if pagePath == "/Operations" {
      pageTitle = "Operations Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Software ~
  // ,  ° . +
  if pagePath == "/Software" {
      pageTitle = "Software Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Events ~
  // ,  ° . +
  if pagePath == "/Events" {
      pageTitle = "Events Page"
      pageList = pageList
  }
  
  
   // ,  ° . +  - _ ` Seating ~
  // ,  ° . +
  if pagePath == "/Seating" {
      pageTitle = "Seating Page"
      pageList = pageList
  }
  
  
   // ,  ° . +  - _ ` Hours ~
  // ,  ° . +
  if pagePath == "/Hours" {
      pageTitle = "Hours Page"
      pageList = pageList
  }


// ,  ° . +
pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  pageHandler


//  .  html url routes 
//  .  as well as terminal cli logs

func main() {

  appName := "~ A Loco Moco Bus // - Website App"


  http.HandleFunc("/", indexHandler)
  
  http.HandleFunc("/user", indexHandler)
  http.HandleFunc("/account", indexHandler)
  http.HandleFunc("/profile", indexHandler)
  
  http.HandleFunc("/portfolio", indexHandler)
  http.HandleFunc("/resume", indexHandler)
  
  http.HandleFunc("/settings", indexHandler)
  
  
  
  Driver
  http.HandleFunc("/Driver", indexHandler)
 + _ Equipment 
 http.HandleFunc("/Equipment", indexHandler)
 + _ Ingredients 
 http.HandleFunc("/Ingredients", indexHandler)
 + _ Vehicle 
 http.HandleFunc("/Vehicle", indexHandler)
 
 
 #### Stuff ~
 + _ Schedule 
 http.HandleFunc("/Schedule", indexHandler)
 + _ Menu
 http.HandleFunc("/Menu", indexHandler)
 + _ Branding 
 http.HandleFunc("/Branding", indexHandler)
 + _ Customers
 http.HandleFunc("/Customers", indexHandler)
 
 #### Stuff ~
 + _ Asana
 http.HandleFunc("/Asana", indexHandler)
 + _ Slack
 http.HandleFunc("/Slack", indexHandler)
 + _ Discord 
 http.HandleFunc("/Discord", indexHandler)
 + _ GumRoad
 http.HandleFunc("/GumRoad", indexHandler)
 
 
#### Stuff ~
+ _ GitHub
http.HandleFunc("/GitHub", indexHandler)
+ _ Google Cloud
http.HandleFunc("/Google_Cloud", indexHandler)
+ _ Linktree
http.HandleFunc("/Linktree", indexHandler)
+ _ Printful
http.HandleFunc("/Printful", indexHandler)

  
  
  
  
  // ,  ° . +` <p>Cafe Needs<ul>
http.HandleFunc("/Menu", indexHandler)
  // ,  ° . + - _ ` Menu ~
  
  http.HandleFunc("/Location", indexHandler)
  // ,  ° . + - _ ` Location ~
  
  http.HandleFunc("/Equiptment", indexHandler)
  // ,  ° . + - _ ` Equiptment ~
  
  http.HandleFunc("/Staff", indexHandler)
 // ,  ° . + - _ ` Staff ~
  
  http.HandleFunc("/Services", indexHandler)
 // ,  ° . + - _ ` Services ~
  
  http.HandleFunc("/Music", indexHandler)
  // ,  ° . +- _ ` Music ~
  
  http.HandleFunc("/Setting", indexHandler)
 // ,  ° . + - _ ` Setting ~
  
  http.HandleFunc("/Sources", indexHandler)
 // ,  ° . + - _ ` Sources ~
  
  http.HandleFunc("/Marketing", indexHandler)
  // ,  ° . +- _ ` Marketing ~
  
  http.HandleFunc("/Operations", indexHandler)
  // ,  ° . +- _ ` Operations ~
  
  http.HandleFunc("/Software", indexHandler)
// ,  ° . +  - _ ` Software ~
  
  http.HandleFunc("/Events", indexHandler)
 // ,  ° . + - _ ` Events ~
  
  http.HandleFunc("/Seating", indexHandler)
  // ,  ° . +- _ ` Seating ~
  
  http.HandleFunc("/Hours", indexHandler)
 // ,  ° . + - _ ` Hours ~



// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
  // ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }


}
