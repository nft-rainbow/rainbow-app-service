# Interface Summery 
## Bind projecter config
This interface can bind the projecter configuration. The `JWT` comes from the NFTRainbow.
### Request Parameters
|Name |Type | Meaning|
|---|---|---|
|guild_id|string|The id of the discord guild|

### Response Parameters
````
"success"
````

### Request Example
`````
curl --request POST
--url https://{DNS_URL}//discord/projecter \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
--data '{
"guild_id": "1008738614827561050"
}'
`````

## Bind activity config
This interface can bind the activity configuration. The `JWT` comes from the NFTRainbow.
### Request Parameters
|Name |Type | Meaning|
|---|---|---|
|contract_id|integer|The id of the contract|
|amount|integer|The amount of the published NFT|
|channel_id|string|The id of the discord channel|
|max_mint_count|integer|The max count of NFT that the user can obtain|
|event|string|The type of the event|
|name|string|The id name of the activity|
|description|string|The description of the activity|

### Response Parameters
````
"success"
````

### Request Example
`````
curl --request POST
--url https://{DNS_URL}/discord/projecter/activity \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
--data '{
    "contract_id":77,
    "amount": 100000,
    "channel_id": "1020536040567615488",
    "max_mint_count": 100,
    "event": "customMint",
    "name": "test",
    "description": "test"
}'
`````

## Get channel information
This interface can bind the activity configuration. The `JWT` comes from the NFTRainbow.

### Response Parameters
The response is the array of the channel struct.

|Name |Type | Meaning|
|---|---|---|
|id|string|The id of the channel|
|guild_id|string|The id of the guild|
|name|string|The name of the channel|
|topic|string|The topic of the channel|
|type|string|The type of the channel|
|last_message_id|string|The id of the last message|
|last_pin_timestamp|time|The timestamp  of the last pin|
|message_count|integer|The count of the message|
|member_count|integer|The count of the member|
|nsfw|bool|Whether the channel is marked as NSFW.|
|icon|string|Icon of the group DM channel|
|position|integer|The position of the channel, used for sorting in client.|
|bitrate|bool|The bitrate of the channel, if it is a voice channel.|
|recipients|[] User|The description of the activity|
|permission_overwrites|[] PermissionOverwrite|The description of the activity|
|user_limit|integer|The user limit of the voice channel.|
|parent_id|integer|The ID of the parent channel, if the channel is under a category. For threads - id of the channel thread was created in|
|rate_limit_per_user|integer|Amount of seconds a user has to wait before sending another message or creating another thread|
|owner_id|integer|The ID of the creator of the group DM or thread|
|application_id|integer|ApplicationID of the DM creator Zeroed if guild channel or not a bot user|
|thread_member|ThreadMember|Thread member object for the current user, if they have joined the thread, only included on certain API endpoints|

The user struct is presented in the following.

|Name |Type | Meaning|
|---|---|---|
|id|string|The id of the channel|
|email|string|The id of the guild|
|username|string|The name of the user|
|avatar|string|The hash of the user's avatar.|
|locale|string|The user's chosen language option|
|discriminator|string|The discriminator of the user|
|token|string|The token of the user. This is only present for the user represented by the current session.|
|verified|bool|Whether the user's email is verified.|
|mfa_enabled|bool|Whether the user has multi-factor authentication enabled|
|banner|string|The hash of the user's banner image|
|accent_color|string|User's banner color, encoded as an integer representation of hexadecimal color code|
|bot|bool|Whether the user is a bot|
|public_flags|string|The public flags on a user's account|
|premium_type|integer|The type of Nitro subscription on a user's account|
|system|bool|Whether the user is an Official Discord System user |
|flags|integer|The flags on a user's account.|


The PermissionOverwrite struct is presented in the following.
|Name |Type | Meaning|
|---|---|---|
|id|string|The id of the channel|
|type|integer|The type of permission_overwrites. 0-PermissionOverwriteTypeRole 1-PermissionOverwriteTypeMember|
|deny|integer|deny permissionOverwrite|
|allow|integer|allow permissionOverwrite|

The ThreadMember struct is presented in the following.
|Name |Type | Meaning|
|---|---|---|
|id|string|The id of the thread|
|user_id|string|The id of the user|
|join_timestamp|time|The timestamp that the user joined the thread|

### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/{guild_id}/channels\
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
`````

### Response Example
``````
[
    {
        "id": "1008738614827561051",
        "guild_id": "1008738614827561050",
        "name": "文字频道",
        "topic": "",
        "type": 4,
        "last_message_id": "",
        "last_pin_timestamp": null,
        "message_count": 0,
        "member_count": 0,
        "nsfw": false,
        "icon": "",
        "position": 0,
        "bitrate": 0,
        "recipients": null,
        "permission_overwrites": [],
        "user_limit": 0,
        "parent_id": "",
        "rate_limit_per_user": 0,
        "owner_id": "",
        "application_id": "",
        "thread_member": null
    },
    {
        "id": "1008738614827561052",
        "guild_id": "1008738614827561050",
        "name": "综合",
        "topic": "",
        "type": 0,
        "last_message_id": "1019809212144029706",
        "last_pin_timestamp": null,
        "message_count": 0,
        "member_count": 0,
        "nsfw": false,
        "icon": "",
        "position": 0,
        "bitrate": 0,
        "recipients": null,
        "permission_overwrites": [
            {
                "id": "1008737738666807416",
                "type": 1,
                "deny": "0",
                "allow": "0"
            }
        ],
        "user_limit": 0,
        "parent_id": "1008738614827561051",
        "rate_limit_per_user": 0,
        "owner_id": "",
        "application_id": "",
        "thread_member": null
    }
]

``````

## Get Project List
This interface can get the list of the project. The `JWT` comes from the NFTRainbow.

### Response Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| count | The number of the uploaded files | integer     |
| items | The files information            | []AdminConfig |

The adminconfig struct is presented in the following.

| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| created_at             | The time of creating the item in the database                        | string     |
| updated_at      | The time of updating the item in the database                          | string     |
| deleted_at   | The time of deleting the item in the database                      | string     |
| id             | The id of the item in the database                        | integer     |
| app_id             | The id of the app                   | integer     |
| guild_id             | The id of the guild                        | string     |
| guild_name             | The name of the guild                       | string     |
| rainbow_user_id             | The id of the user                       | integer     |


### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/projecter/ \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
`````

### Response Example
```
{
    "count": 1,
    "items": [
        {
            "id": 1,
            "created_at": "2022-10-19T15:07:25.947+08:00",
            "updated_at": "2022-10-19T15:07:25.947+08:00",
            "deleted_at": null,
            "app_id": 2,
            "guild_id": "1008738614827561050",
            "guild_name": "xqYang的服务器",
            "rainbow_user_id": 2
        }
    ]
}
```

## Get project information
This interface can get the detail information of the project according to the id. The `JWT` comes from the NFTRainbow.

### Response Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| created_at             | The time of creating the item in the database                        | string     |
| updated_at      | The time of updating the item in the database                          | string     |
| deleted_at   | The time of deleting the item in the database                      | string     |
| id             | The id of the item in the database                        | integer     |
| app_id             | The id of the app                   | integer     |
| guild_id             | The id of the guild                        | string     |
| guild_name             | The name of the guild                       | string     |
| rainbow_user_id             | The id of the user                       | integer     |


### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/projecter/{id} \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
`````

### Response Example
```
{
    "id": 1,
    "created_at": "2022-10-19T15:07:25.947+08:00",
    "updated_at": "2022-10-19T15:07:25.947+08:00",
    "deleted_at": null,
    "app_id": 2,
    "guild_id": "1008738614827561050",
    "guild_name": "xqYang的服务器",
    "rainbow_user_id": 2
}
```

## Get Activity List
This interface can get the list of the activity. The `JWT` comes from the NFTRainbow.

### Response Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| count | The number of the uploaded files | integer     |
| items | The files information            | []ActivityConfig |

The activityConfig struct is presented in the following.

| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| created_at             | The time of creating the item in the database                        | string     |
| updated_at      | The time of updating the item in the database                          | string     |
| deleted_at   | The time of deleting the item in the database                      | string     |
| id             | The id of the item in the database                        | integer     |
| app_id             | The id of the app                   | integer     |
|contract_id|integer|The id of the contract|
|amount|integer|The amount of the published NFT|
|channel_id|string|The id of the discord channel|
|max_mint_count|integer|The max count of NFT that the user can obtain|
|event|string|The type of the event|
|name|string|The id name of the activity|
|description|string|The description of the activity|
|contract_type|string|The type of the contract|
|contract_address|string|The address of the contract|
|chain_type|integer|The type of the chain|


### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/projecter/activity \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
`````

### Response Example
```
{
    "count": 2,
    "items": [
        {
            "id": 1,
            "created_at": "2022-10-19T16:27:06.307+08:00",
            "updated_at": "2022-10-19T16:27:06.307+08:00",
            "deleted_at": null,
            "contract_id": 75,
            "channel_id": "1020536040567615488",
            "amount": 100000,
            "max_mint_count": 100,
            "event": "customMint",
            "name": "test",
            "description": "test",
            "app_id": 2,
            "contract_type": 1,
            "contract_address": "cfxtest:acb4xr6k3dr01k5fvddeevyz9abc5cvkja3cfsstnt",
            "chain_type": 1
        },
        {
            "id": 2,
            "created_at": "2022-10-19T16:46:47.14+08:00",
            "updated_at": "2022-10-19T16:46:47.14+08:00",
            "deleted_at": null,
            "contract_id": 77,
            "channel_id": "1020536040567615488",
            "amount": 100000,
            "max_mint_count": 100,
            "event": "customMint",
            "name": "test",
            "description": "test",
            "app_id": 2,
            "contract_type": 1,
            "contract_address": "cfxtest:achj1vcsns2w8kvmdctxrwwunzvmekcfayf4cv2hse",
            "chain_type": 1
        }
    ]
}
```

## Get Activity information
This interface can get the detail information of the activity. The `JWT` comes from the NFTRainbow.

### Response Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| created_at             | The time of creating the item in the database                        | string     |
| updated_at      | The time of updating the item in the database                          | string     |
| deleted_at   | The time of deleting the item in the database                      | string     |
| id             | The id of the item in the database                        | integer     |
| app_id             | The id of the app                   | integer     |
|contract_id|integer|The id of the contract|
|amount|integer|The amount of the published NFT|
|channel_id|string|The id of the discord channel|
|max_mint_count|integer|The max count of NFT that the user can obtain|
|event|string|The type of the event|
|name|string|The id name of the activity|
|description|string|The description of the activity|
|contract_type|string|The type of the contract|
|contract_address|string|The address of the contract|
|chain_type|integer|The type of the chain|


### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/projecter/activity/{id} \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
`````

### Response Example
```
{
    "id": 1,
    "created_at": "2022-10-19T16:27:06.307+08:00",
    "updated_at": "2022-10-19T16:27:06.307+08:00",
    "deleted_at": null,
    "contract_id": 75,
    "channel_id": "1020536040567615488",
    "amount": 100000,
    "max_mint_count": 100,
    "event": "customMint",
    "name": "test",
    "description": "test",
    "app_id": 2,
    "contract_type": 1,
    "contract_address": "cfxtest:acb4xr6k3dr01k5fvddeevyz9abc5cvkja3cfsstnt",
    "chain_type": 1
}
```


### Discord Login 
This interface can obtain the JWT.

### Request Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| user_id             | The id of user                       | string     |
| channel_id      | The id of channel                          | string     |

### Response Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| expire             | The time of creating the item in the database                        | string     |
| token      | The time of updating the item in the database                          | string     |

### Request Example
`````
curl --request POST
--url https://{DNS_URL}/discord/login \
--header 'Content-Type: application/json' \
`````

### Response Example
```
{
    "expire": "2022-10-20T17:47:12.9506734+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjYyNTkyMzIsIm9yaWdfaWF0IjoxNjY2MjU1NjMyfQ.-K-FfZfU7vx5ZH4Fjw2MrwaCDeQuwOWEW8v_NJJls3k"
}
```

## Bind User Address
This interface can bind the cfx address with the discord_user_id. The `JWT` comes from the `Discord Login` interface.

### Request Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| user_id             | The id of user                       | string     |
| user_address      | The address of user                          | string     |

### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/projecter/activity/{id} \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
--data `{
    "user_id": "830409486455668766",
    "user_address": "cfxtest:aajb342mw5kzad6pjjkdz0wxx0tr54nfwpbu6yaj49"
}`
`````

### Response Example
```
"success"
```

## Custom Mint
This interface can execute the customMint. The `JWT` comes from the `Discord Login` interface.

### Request Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| user_id             | The id of user                       | string     |
| channel_id      | The id of channel                          | string     |

### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/user/mint \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
--data `{
    "user_id": "830409486455668766",
    "channel_id": "1020536040567615488"
}`
`````

### Response Example
```
"success"
```

## Get Binding address
This interface can get the binding address. The `JWT` comes from the `Discord Login` interface.

### Response Parameters
| Name  | Meaning                          | Type        |
| ----- | -------------------------------- | ----------- |
| created_at             | The time of creating the item in the database                        | string     |
| updated_at      | The time of updating the item in the database                          | string     |
| deleted_at   | The time of deleting the item in the database                      | string     |
| id             | The id of the item in the database                        | integer     |
| user_id             | The id of user                       | string     |
| user_address      | The address of user                          | string     |


### Request Example
`````
curl --request GET
--url https://{DNS_URL}/discord/user/{user_id} \
--header 'Authorization: Bearer {JWT}' \
--header 'Content-Type: application/json' \
`````

### Response Example
```
{
    "id": 1,
    "created_at": "2022-10-20T16:48:49.271+08:00",
    "updated_at": "2022-10-20T16:48:49.271+08:00",
    "deleted_at": null,
    "user_id": "830409486455668766",
    "user_address": "cfxtest:aajb342mw5kzad6pjjkdz0wxx0tr54nfwpbu6yaj49"
}
```


