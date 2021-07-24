print('Start #################################################################');

db.createUser(
    {
        user: "arualcms",
        pwd: "SDF678FG",
        roles: [
            {
                role: "readWrite",
                db: "arualcms"
            }
        ]
    }
);
db = new Mongo().getDB("arualcms");

db.createCollection('languages');
db.createCollection('posts');
db.createCollection('settings');
db.createCollection('storage');
db.createCollection('texts');
db.createCollection('users');


db.languages.insert(
    [
        {
            "key": "en",
            "value": "English",
            "default": 1
        },
        {
            "key": "cz",
            "value": "Čeština",
            "default": 0
        }
    ]
);

db.posts.insert([
    {
        "title": {
            "en": "Post 1",
            "cz": "CZ Post 1"
        },
        "body": {
            "en": "<p>asdsad<\/p><p><strong>asdsad<\/strong><\/p>",
            "cz": "<p>CZ asdsad<\/p><p><strong>asdsad<\/strong><\/p>"
        },
        "excerpt": {
            "en": "Excerpt 1",
            "cz": "CZ Excerpt 1"
        },
        "published": "2021-03-02T10:00:00.000Z",
        "meta": {
            "title": {
                "en": "Post 1",
                "cz": "CZ Post 1"
            },
            "keywords": {
                "en": "mock keywords 1",
                "cz": "mock keywords 1"
            },
            "description": {
                "en": "Mock post 1",
                "cz": "CZ Mock post 1"
            }
        },
        "id": 1,
        "file": "154154879_354719135650311_7324051498172828998_n-2.png",
        "src": "http:\/\/localhost:8000\/storage\/154154879_354719135650311_7324051498172828998_n-2.png"
    },
    {
        "title": {
            "en": "Post 2",
            "cz": "CZ Post 2"
        },
        "body": {
            "en": "<p>asdsad<\/p><p><strong>asdsad<\/strong><\/p>",
            "cz": "<p>CZ asdsad<\/p><p><strong>asdsad<\/strong><\/p>"
        },
        "excerpt": {
            "en": "Excerpt 2",
            "cz": "CZ Excerpt 2"
        },
        "published": "2021-03-02T10:00:00.000Z",
        "meta": {
            "title": {
                "en": "Post 2",
                "cz": "CZ Post 2"
            },
            "keywords": {
                "en": "mock keywords 2",
                "cz": "mock keywords 2"
            },
            "description": {
                "en": "Mock post 2",
                "cz": "CZ Mock post 2"
            }
        },
        "id": 2,
        "file": "154154879_354719135650311_7324051498172828998_n-2.png",
        "src": "http:\/\/localhost:8000\/storage\/154154879_354719135650311_7324051498172828998_n-2.png"
    },
    {
        "title": {
            "en": "Post 3",
            "cz": "CZ Post 3"
        },
        "body": {
            "en": "<p>asdsad<\/p><p><strong>asdsad<\/strong><\/p>",
            "cz": "<p>CZ asdsad<\/p><p><strong>asdsad<\/strong><\/p>"
        },
        "excerpt": {
            "en": "Excerpt 3",
            "cz": "CZ Excerpt 3"
        },
        "published": "2021-03-02T10:00:00.000Z",
        "meta": {
            "title": {
                "en": "Post 3",
                "cz": "CZ Post 3"
            },
            "keywords": {
                "en": "mock keywords 3",
                "cz": "mock keywords 3"
            },
            "description": {
                "en": "Mock post 3",
                "cz": "CZ Mock post 3"
            }
        },
        "id": 3,
        "file": "154154879_354719135650311_7324051498172828998_n-2.png",
        "src": "http:\/\/localhost:8000\/storage\/154154879_354719135650311_7324051498172828998_n-2.png"
    }
]);

db.settings.insert([
    {
        "key": "title",
        "value": {
            "en": "arualCMS",
            "cz": "CZ arualCMS"
        }
    },
    {
        "key": "subtitle",
        "value": {
            "en": "Simpliest CMS for your need.",
            "cz": "Simpliest CMS for your need."
        }
    },
    {
        "key": "author",
        "value": {
            "en": "Ale\u0161 Jandera, shopyCRM",
            "cz": "Ale\u0161 Jandera, shopyCRM"
        }
    },
    {
        "key": "contactEmail",
        "value": {
            "en": "jandera@shopycrm.com",
            "cz": "jandera@shopycrm.com"
        }
    },
    {
        "key": "contactPhone",
        "value": {
            "en": "+421 904 750 220",
            "cz": "+420 904 750 220"
        }
    },
    {
        "key": "metaTitle",
        "value": {
            "en": "arualCMS simpliest cms for your need.",
            "cz": "CZ arualCMS simpliest cms for your need."
        }
    },
    {
        "key": "metaDescription",
        "value": {
            "en": "arualCMS simpliest cms for your need.",
            "cz": "CZ arualCMS simpliest cms for your need."
        }
    },
    {
        "key": "metaKeywords",
        "value": {
            "en": "cms, arualcms, content management system",
            "cz": "cms, arualcms, content management system"
        }
    },
    {
        "key": "smtp",
        "value": {
            "en": "smtp.gmail.com",
            "cz": "smtp.gmail.com"
        }
    },
    {
        "key": "smtp_user",
        "value": {
            "en": "arualcms@ajandera.com",
            "cz": "arualcms@ajandera.cz"
        }
    },
    {
        "key": "smtp_password",
        "value": {
            "en": "password",
            "cz": "CZ password"
        }
    },
    {
        "key": "smtp_port",
        "value": {
            "en": "567",
            "cz": "CZ 567"
        }
    }
]);

db.texts.insert([
    {
        "key": "what_we_do",
        "value": {
            "en": "[en] Lorem ipsum dolor sit amet, consectetur adipisicing elit. A deserunt neque tempore recusandae animi soluta quasi? Asperiores rem dolore eaque vel, porro, soluta unde debitis aliquam laboriosam. Repellat explicabo, maiores!\n\nLorem ipsum dolor sit amet, consectetur adipisicing elit. Omnis optio neque consectetur consequatur magni in nisi, natus beatae quidem quam odit commodi ducimus totam eum, alias, adipisci nesciunt voluptate. Voluptatum.",
            "cz": "[CZ] Lorem ipsum dolor sit amet, consectetur adipisicing elit. A deserunt neque tempore recusandae animi soluta quasi? Asperiores rem dolore eaque vel, porro, soluta unde debitis aliquam laboriosam. Repellat explicabo, maiores!\n\nLorem ipsum dolor sit amet, consectetur adipisicing elit. Omnis optio neque consectetur consequatur magni in nisi, natus beatae quidem quam odit commodi ducimus totam eum, alias, adipisci nesciunt voluptate. Voluptatum."
        }
    },
    {
        "key": "contact",
        "value": {
            "en": "<p>Start Bootstrap&nbsp;<\/p><p>3481 Melrose Place&nbsp;<\/p><p>Beverly Hills, CA 90210&nbsp;<\/p><p>P:&nbsp;(123) 456-7890&nbsp;<\/p><p>E:&nbsp;<a href=\"mailto:#\" target=\"_blank\" style=\"color: rgb(0, 123, 255);\">name@example.com<\/a><\/p>",
            "cz": "<p>[CZ] Start Bootstrap&nbsp;<\/p><p>3481 Melrose Place&nbsp;<\/p><p>Beverly Hills, CA 90210&nbsp;<\/p><p>P:&nbsp;(123) 456-7890&nbsp;<\/p><p>E:&nbsp;<a href=\"mailto:#\" target=\"_blank\" style=\"color: rgb(0, 123, 255);\">name@example.com<\/a><\/p>"
        }
    }
]);

db.users.insert([
        {
            "username": "arualcms@arualcms.com",
            "password": "$2y$10$U2KiZd7FZVLDCgULweu\/mOYHqe0yJHOuYK2jTjsy4oA825GWYr5pS",
            "id": 1
        }
    ]
);

print('END #################################################################');
