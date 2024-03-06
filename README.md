#  Golden Hands of Beauty :nail_care:
Our code is a Go application for Beauty salon web service. It implements HTTP request handlers, work with MySQL database, sessions, as well as functions for error handling and HTML templating

 ## Key features:
 Registration and Authentication of Users:
>User registration form with mandatory fields: name, email, password and phone number.
Data validity check and adding a new user to the database.
Login form with user authentication by email and password.
>>Authorisation and Session Management:
Set and delete sessions for user authentication.
Protection against CSRF attacks using tokens in forms.
>>>Displaying Different Pages:
Homepage with user information.
Spa, cosmetology, grooming, about us, booking and news pages.
>>>>Processing of Booking Requests:
Displaying the booking page with user details.
Ability to book services and products.
>>>>>Display of News and Cosmetic Products:
Displaying the latest news and cosmetic products on their respective pages.
>>>>>>Error Logging and Handling:
Logging information about each HTTP request to a log.
Error handling, including panic recovery and sending appropriate HTTP responses.
>>>>>>>Header Security and Protection:
Setting secure HTTP headers to protect against XSS and clickjacking attacks.
Cross-site request forgery (CSRF) protection.
>>>>>>>>Splitting into Pages, Layouts, and Partial Templates:
Using templates to split into pages, layouts, and partial templates, making it easier to maintain and extend code.
>>>>>>>>>Logging Information and Errors:
Logging information about application actions and errors for later analysis.
Installing and Using Template Functions:
Creating and using a template cache with a function to format a date into a legible form.
>>>>>>>>>>>Handling Panic States:
Handling panic states to prevent application crashing and logging related errors.

### Main codes:
>Hadler.go
This code includes methods for rendering HTML page templates and processing form data such as user registration, authentication, logout and displaying information about beauty salon products and services

![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/18f4e620-6f3f-40a8-b571-aa400ec86429)
![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/d5e4b93f-707f-41bc-8e7c-d513bf2a058c)

>>helpers.go
This code snippet is a set of methods for error handling, user authentication, and rendering web application templates.

![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/2e4da55f-7672-4c2f-8ff7-8a9bb602c090)
![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/e206dab0-3640-4a1e-ba7e-64a48d13a906)

>>>middleware.go
This code fragment represents a set of middleware functions for processing HTTP requests from a web application. Middleware is software that sits between the client and the main application and performs additional operations on requests and responses.

![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/43e73181-d8b1-49af-86af-fc50186db295)
![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/fa06354e-00fb-48de-8644-7b90c57b3ef4)

>>>>routes.go is the definition of routes for a web application. It uses the pat and alice packages to handle routes and middleware.

![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/c3d84011-fc75-4489-9bf7-20976696d711)


>>>>>templates.go It contains functions and structures that handle the HTML templates used to display web pages.
1. The `humanDate` function defines a custom function to format the time in a human-readable format.

2. `functions` is a global variable that contains custom template functions. In this case, it contains only one function, `humanDate`.

3. `newTemplateCache` is a function that creates a template cache based on template files located in the specified directory. It uses the filepath and text/template packages to handle the file system and templates.

4. `templateData` is a data structure that contains information that will be passed to HTML templates for display on web pages. It contains various fields such as CSRF token, current year, form, flash message, user authentication information, etc.

   ![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/03ff52b5-dcef-4bd6-8437-71e8e6fd2c76)
   ![image](https://github.com/hvruu/AdvancedProgAssingment/assets/147140948/bae537e8-7a0a-4efe-aaf7-0b750ee92df1)

   :notebook: :black_nib: **Authors:** 
   **Abai Akylzhanuly,Bekzat Stanbek,Galymzhan Syrymbet** 
**Github link: https://github.com/hvruu/AdvancedProgAssingment** :shipit:



