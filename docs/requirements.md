# Server
## Setup
- Docker
- Server
- Databases
## Routes
- Base
    - Domain: (www.) brypt.com
    - Base
        - Description:
        - Corequisites: NA
    - About
        - Description:
        - Corequisites: NA
    - Contact
        - Description:
        - Corequisites: NA
    - Privacy Policy
        - Description:
        - Corequisites: NA
    - Terms of Service
        - Description:
        - Corequisites: NA
    - Binaries Download
        - Description:
        - Corequisites: NA
    - Embedded Devices
        - Description:
        - Corequisites: NA
- Access
    - Domain: access.brypt.com
    - Register
        - Description:
        - Corequisites: Users Table, Authentication Method
    - Login
        - Description:
        - Corequisites: Users Table, Authentication Method
    - Delete
        - Description:
        - Corequisites: Users Table, Authentication Method
- Bridge
    - Domain: bridge.brypt.com
    - Register
        - Description:
        - Corequisites: Networks Document, Nodes Table, Users Table
    - Join
        - Description:
        - Corequisites: Networks Document, Nodes Table, Users Table
    - Delete
        - Description:
        - Corequisites: Networks Document, Nodes Table, Users Table
- Users
    - Domain: (www.)brypt.com
    - Profile
        - Description:
        - Corequisites: Users Table, Authentication Method
    - Settings
        - Description:
        - Corequisites: Users Table, Authentication Method
- Dashboard
    -Domain: dashboard.brypt.com
    - Base:
        - Description:
        -Corequisites: Networks Table, Users Table, Authentication Method

## Authentication
- JWT or OAuth or Nah

# Database
## Users
- MySQL
    - id
        - Primary Key
        - 32 Char
    -username
        - 32 Char
    -first_name
        - 32 Char
    -last_name
        - 255 Char
    -email
        - 255 Char
    -organization
        - 255 Char
    -networks
        - Array
            - 32 Char
    -age
        - Datetime
    -join_date
        - Datetime
    -last_login
        - Datetime
    -login_attempts
        - Integer
    -login_token
        - 255 Char
    -region
        - 255 Char

## Networks
- MongoDB
    -id
        - Primary Key
        - String
    -network_name
        - String
    -owner_name
        - Foreign Key
        - String
    -managers
        - Array
            - Foreign Key
            - String
    -direct_peers
        - Integer
    - total_peers
        - Integer
    - IP_address
        - String
    - port
        - Integer
    - connection_token
        - String
    - clusters
        -JSON Array
           - id
              - String
           - connection_token
              - String
           - coord_IP
              - String
           - coord_port
              - String
           - comm_tech
              - String
    - created_on
        - Datetime
    - last_accessed
        - Datetime

## Nodes
- MySQL
    - id
        - Primary Key
        - 32 Char
    - serial_number
        - 255 Char
    - type
        - 32 Char
    - created_on
        - Datetime
    - registered_on
        - Datetime
    - registered_to
        - Foreign Key
        - 32 Char
    - connected_network
        - Foreign Key
        - 32 Char

# Routing

# Nodes

# Security Protocol
