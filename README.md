![code coverage badge](https://github.com/Janisgee/discovery_app/actions/workflows/ci.yml/badge.svg)

# Discovery App

![Discovery-app-logo-2](https://github.com/user-attachments/assets/4e920fa3-1e17-429a-b5f3-cf844f59433f)

Discovery App is an application helping user to get place knowladge around the world. Using Next.js to build its front-end interface and Go framework to create back-end server. Using AI models generated places contents. And using Postgres SQL to store user's bookmarks and accounts data.

https://github.com/user-attachments/assets/1a7bcee6-9dcf-45eb-9778-d17f2fe9c536

Discovery helps user to explore new places to visit. By storing user's preference, user able to revisit the bookmarks for further research.

To get started, see the docs below and resources.

## Local Development

1. Make sure you're on Go version 1.22+.

    Run the server to install all the packages for this application:
    
    ```
    /scripts/build.sh
    ```

2. Create environment file (.env.local) inside discovery-web for front-end server to use. (Fill in your own key value according different variables)
     ```
    # Front-end Environment
    
    # Cloudinary (profile picture storage) Configure for front-end environment
    NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME=""
    NEXT_PUBLIC_CLOUDINARY_API_KEY=""
    CLOUDINARY_UPLOAD_PRESET=""
    CLOUDINARY_API_SECRET=""
    
    # Time limit for autologout if user is not active
    NEXT_PUBLIC_AUTOLOGOUT_TIME=1800000
    
    # Server Base URL
    NEXT_PUBLIC_API_SERVER_BASE="http://localhost:8080"
    ```

3.  Create environment file (.env) inside discovery-api for back-end server to use. (Fill in your own key value according different variables)

     ```
    # Back-end Environment

    # Indicate whether your application is running in a local or production environment (such as on Render). "" is from local.
    IS_RENDER= ""
   
    # Connect to ChatGPT API
    OPENAI_API_KEY = ''
    
    # Connect to port
    PORT = 8080
    
    # PostgreSQL connection string - connect to database (discovery)
    PSQL_CONNECTION_STRING = "postgres://postgres:postgres@localhost:5433/discovery_app?sslmode=disable"
    
    # Connect to docker database name: discovery_app
    # Generate PostgresSQL request
    #goose postgres "postgres://postgres:postgres@localhost:5433/discovery_app" up
    
    # google mailer - connect to gmail to send user password reset email
    GGMAILER_KEY=""
    GGMAILER_EMAIL=""
    
    # Google map - get google placeID
    GMAPS_API_KEY=''
    
    # pexels image - get location image
    PEXEL_API_KEY=''
    
    # Client Url
    CLIENT_BASE_URL = "http://localhost:3000"
     ```
4. Start PostgresSQL database
   ```
   /scripts/start_database.sh
   ```


5. Generate database code
   ```
   /scripts/generate_db_code.sh
   ```



## Quick Start

- Start running front-end server on (http://localhost:3000/)

  ```
  pnpm dev
  ```

- Start running back-end server on (http://localhost:8080/)
  - Build app server and run.
   ```
   /scripts/build_run.sh
   ```

- Database schema migration with "goose up" : Updating its schema to the latest version (cd sql/schema directory)
  ```
   goose postgres "postgres://postgres:postgres@localhost:5433/discovery_app" up
  ```

- Database schema migration with "goose down" : Revert the changes to the previous version (cd sql/schema directory)
  ```
   goose postgres "postgres://postgres:postgres@localhost:5433/discovery_app" down
  ```


