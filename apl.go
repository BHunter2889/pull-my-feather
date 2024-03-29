package main

const (
	aplJson = `{
    "document": {
        "type": "APL",
        "version": "1.1",
        "theme": "dark",
        "import": [
            {
                "name": "alexa-layouts",
                "version": "1.0.0"
            }
        ],
        "settings": {},
        "resources": [
            {
                "description": "Stock color for the light theme",
                "colors": {
                    "colorTextPrimary": "#565A5D"
                }
            },
            {
                "description": "Stock color for the dark theme",
                "when": "${viewport.theme == 'dark'}",
                "colors": {
                    "colorTextPrimary": "#DDDDDD"
                }
            },
            {
                "description": "Standard font sizes",
                "dimensions": {
                    "textSizeBody": 48,
                    "textSizeBodySecondary": 36,
                    "textSizePrimary": 30,
                    "textSizeSecondary": 23,
                    "textSizeSecondaryHint": 25
                }
            },
            {
                "description": "Common spacing values",
                "dimensions": {
                    "spacingThin": 6,
                    "spacingSmall": 12,
                    "spacingMedium": 24,
                    "spacingLarge": 48,
                    "spacingExtraLarge": 72
                }
            },
            {
                "description": "Common margins and padding",
                "dimensions": {
                    "marginTop": 40,
                    "marginLeft": 60,
                    "marginRight": 60,
                    "marginBottom": 40,
					"paddingLeft": 30,
					"paddingLeft40": 40,
					"paddingRight": 30,
					"paddingRight72": 72,
					"paddingBottom": 200,		
					"paddingTop40": 40,
					"paddingTop50": 50,
					"primaryTextPaddingTop": 35,
					"bulletPointPaddingTop": 10
                }
            }
        ],
        "styles": {
            "textStyleBase": {
                "description": "Base font description; set color",
                "values": [
                    {
                        "color": "@colorTextPrimary"
                    }
                ]
            },
            "textStyleBase0": {
                "description": "Thin version of basic font",
                "extend": "textStyleBase",
                "values": {
                    "fontWeight": "100"
                }
            },
            "textStyleBase1": {
                "description": "Light version of basic font",
                "extend": "textStyleBase",
                "values": {
                    "fontWeight": "300"
                }
            },
            "mixinBody": {
                "values": {
                    "fontSize": "@textSizeBody"
                }
            },
            "mixinBodySecondary": {
                "values": {
                    "fontSize": "@textSizeBodySecondary",
                    "fontWeight": 400
                }
            },
            "mixinPrimary": {
                "values": {
                    "fontSize": "@textSizePrimary"
                }
            },
            "mixinSecondary": {
                "values": {
                    "fontSize": "@textSizeSecondary"
                }
            },
            "textStylePrimary": {
                "extend": [
                    "textStyleBase1",
                    "mixinPrimary"
                ]
            },
            "textStyleSecondary": {
                "extend": [
                    "textStyleBase0",
                    "mixinSecondary"
                ]
            },
            "textStyleBody": {
                "extend": [
                    "textStyleBase1",
                    "mixinBody"
                ]
            },
            "textStyleBodySecondary": {
                "extend": [
                    "textStyleBase1",
                    "mixinBodySecondary"
                ]
            },
            "textStyleSecondaryHint": {
                "values": {
                    "fontFamily": "Bookerly",
                    "fontStyle": "italic",
                    "fontSize": "@textSizeSecondaryHint",
                    "color": "@colorTextPrimary"
                }
            }
        },
        "onMount": [],
        "graphics": {},
        "commands": {},
        "layouts": {},
        "mainTemplate": {
            "parameters": [
                "payload"
            ],
            "items": [
                {
                    "when": "${viewport.shape == 'round'}",
                    "type": "Container",
                    "direction": "column",
                    "width": "100vw",
                    "height": "100vh",
                    "items": [
                        {
                            "type": "Image",
                            "source": "${payload.bodyTemplateData.image.sources[0].url}",
                            "scale": "best-fill",
                            "width": "100vw",
                            "height": "100vh",
                            "position": "absolute",
                            "overlayColor": "rgba(0, 0, 0, 0.6)"
                        },
                        {
                            "type": "ScrollView",
                            "width": "100vw",
                            "height": "100vh",
                            "item": [
                                {
                                    "type": "Container",
                                    "direction": "column",
                                    "alignItems": "center",
                                    "paddingLeft": "@paddingLeft",
                                    "paddingRight": "@paddingRight",
                                    "paddingBottom": "@paddingBottom",
                                    "items": [
                                        {
                                            "type": "AlexaHeader",
                                            "headerAttributionImage": "${payload.bodyTemplateData.logoUrl}",
                                            "headerTitle": "${payload.bodyTemplateData.title}"
                                        },
                                        {
                                            "type": "Text",
                                            "text": "<b>${payload.bodyTemplateData.textContent.subheader.text}</b>",
                                            "style": "textStylePrimary",
                                            "color": "#62A0A5",
                                            "width": "90vw",
                                            "textAlign": "center"
                                        },
                                        {
                                            "type": "Text",
                                            "text": "<b>${payload.bodyTemplateData.textContent.title.text}</b>",
                                            "style": "textStyleBody",
                                            "width": "90vw",
                                            "textAlign": "center"
                                        },
                                        {
                                            "type": "Text",
                                            "text": "${payload.bodyTemplateData.textContent.subtitle.text}",
                                            "style": "textStylePrimary",
                                            "width": "90vw",
                                            "textAlign": "center"
                                        },
                                        {
                                            "type": "Text",
                                            "text": "${payload.bodyTemplateData.textContent.primaryText.text}",
                                            "paddingTop": "@paddingTop40",
                                            "style": "textStylePrimary",
                                            "width": "90vw",
                                            "textAlign": "center"
                                        },
                                        {
                                            "type": "Text",
                                            "text": "${payload.bodyTemplateData.textContent.bulletPoints[0].text}",
                                            "paddingTop": "@paddingTop50",
                                            "style": "textStylePrimary",
                                            "width": "90vw",
                                            "textAlign": "center"
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                },
                {
                    "type": "Container",
                    "width": "100vw",
                    "height": "100vh",
                    "items": [
                        {
                            "type": "Image",
                            "source": "${payload.bodyTemplateData.backgroundImage.sources[0].url}",
                            "scale": "best-fill",
                            "width": "100vw",
                            "height": "100vh",
                            "position": "absolute"
                        },
                        {
                            "type": "AlexaHeader",
                            "headerTitle": "${payload.bodyTemplateData.title}",
                            "headerAttributionImage": "${payload.bodyTemplateData.logoUrl}"
                        },
                        {
                            "type": "Container",
                            "direction": "row",
                            "paddingLeft": "@paddingLeft40",
                            "paddingRight": "@paddingRight72",
                            "grow": 1,
                            "items": [
                                {
                                    "type": "Image",
                                    "source": "${payload.bodyTemplateData.image.sources[0].url}",
                                    "width": "35vw",
                                    "height": "65vh",
                                    "scale": "best-fit",
                                    "align": "center"
                                },
                                {
                                    "type": "ScrollView",
                                    "height": "60vh",
                                    "shrink": 1,
                                    "item": [
                                        {
                                            "type": "Container",
                                            "items": [
                                                {
                                                    "type": "Text",
                                                    "text": "<b>${payload.bodyTemplateData.textContent.subheader.text}</b>",
                                                    "style": "textStylePrimary",
                                                    "color": "#62A0A5"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "<b>${payload.bodyTemplateData.textContent.title.text}</b>",
                                                    "style": "textStyleBody"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "${payload.bodyTemplateData.textContent.subtitle.text}",
                                                    "style": "textStylePrimary"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "${payload.bodyTemplateData.textContent.primaryText.text}",
                                                    "paddingTop": "@primaryTextPaddingTop",
													"paddingRight": "@paddingRight",
                                                    "style": "textStyleBodySecondary"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "${payload.bodyTemplateData.textContent.bulletPoints[0].text}",
                                                    "paddingTop": "@bulletPointPaddingTop",
                                                    "paddingRight": "@paddingRight",
                                                    "style": "textStylePrimary"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                },
				{
					"when": "${@viewportProfile == @tvLandscapeXLarge}",
                    "type": "Container",
                    "width": "100vw",
                    "height": "100vh",
                    "items": [
                        {
                            "type": "Image",
                            "source": "${payload.bodyTemplateData.backgroundImage.sources[0].url}",
                            "scale": "best-fill",
                            "width": "100vw",
                            "height": "100vh",
                            "position": "absolute"
                        },
                        {
                            "type": "AlexaHeader",
                            "headerTitle": "${payload.bodyTemplateData.title}",
                            "headerAttributionImage": "${payload.bodyTemplateData.logoUrl}"
                        },
                        {
                            "type": "Container",
                            "direction": "row",
                            "paddingLeft": "@paddingLeft40",
                            "paddingRight": "@paddingRight72",
                            "grow": 1,
                            "items": [
                                {
                                    "type": "Image",
                                    "source": "${payload.bodyTemplateData.image.sources[0].url}",
                                    "width": "35vw",
                                    "height": "65vh",
                                    "scale": "best-fit",
                                    "align": "center"
                                },
                                {
                                    "type": "ScrollView",
                                    "height": "60vh",
                                    "shrink": 1,
                                    "item": [
                                        {
                                            "type": "Container",
                                            "items": [
                                                {
                                                    "type": "Text",
                                                    "text": "<b>${payload.bodyTemplateData.textContent.subheader.text}</b>",
                                                    "style": "textStyleSecondaryHint",
                                                    "color": "#62A0A5"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "<b>${payload.bodyTemplateData.textContent.title.text}</b>",
                                                    "style": "textStyleBodySecondary"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "${payload.bodyTemplateData.textContent.subtitle.text}",
                                                    "style": "textStyleSecondaryHint"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "${payload.bodyTemplateData.textContent.primaryText.text}",
                                                    "paddingTop": "@primaryTextPaddingTop",
													"paddingRight": "@paddingRight",
                                                    "style": "textStylePrimary"
                                                },
                                                {
                                                    "type": "Text",
                                                    "text": "${payload.bodyTemplateData.textContent.bulletPoints[0].text}",
                                                    "paddingTop": "@bulletPointPaddingTop",
                                                    "paddingRight": "@paddingRight",
                                                    "style": "textStyleSecondary"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    },
    "datasources": {
        "bodyTemplateData": {
            "type": "object",
            "objectId": "bt3Sample",
            "backgroundImage": {
                "contentDescription": null,
                "smallSourceUrl": null,
                "largeSourceUrl": null,
                "sources": [
                    {
                        "url": "https://1904upload.s3.amazonaws.com/1904diamonds_EDIT.png",
                        "size": "small",
                        "widthPixels": 0,
                        "heightPixels": 0
                    },
                    {
                        "url": "https://1904upload.s3.amazonaws.com/1904diamonds_EDIT.png",
                        "size": "large",
                        "widthPixels": 0,
                        "heightPixels": 0
                    }
                ]
            },
            "title": "Tickled Pink Award Winner",
            "image": {
                "contentDescription": null,
                "smallSourceUrl": null,
                "largeSourceUrl": null,
                "sources": [
                    {
                        "url": "https://1904upload.s3.amazonaws.com/PinkFeather.png",
                        "size": "small",
                        "widthPixels": 0,
                        "heightPixels": 0
                    },
                    {
                        "url": "https://1904upload.s3.amazonaws.com/PinkFeather.png",
                        "size": "large",
                        "widthPixels": 0,
                        "heightPixels": 0
                    }
                ]
            },
            "textContent": {
                "subheader": {
                    "type": "Text",
                    "text": "Monday, August 26, 2019"
                },
				"title": {
                    "type": "Text",
                    "text": "Daniel Dorsey"
                },
                "subtitle": {
                    "type": "Text",
                    "text": "From: Laura Tromben"
                },
                "primaryText": {
                    "type": "Text",
                    "text": " Praise "
                },
                "bulletPoints": [
					{
                    	"type": "Text",
                    	"text": "• Being super patient and helping me with Code with Pride "
                	}
				]
            },
            "logoUrl": "https://1904upload.s3.amazonaws.com/1904labsNoTag_darkBG.png",
            "hintText": "Try, \"Alexa, tickle me pink\""
        }
    }
}`
)
