const path = require('path')
require("dotenv").config()
module.exports = {
    env:{
        API_URL: process.env.API_URL,
        NEXT_PUBLIC_GOOGLE_ANALYTICS: process.env.NEXT_PUBLIC_GOOGLE_ANALYTICS
    },
    future: {
      webpack5: true,
    },
    webpack:config =>{
        config.resolve.alias['components']= path.join(__dirname,'components')
        config.resolve.alias['public']= path.join(__dirname,'public')
        config.resolve.alias['styles']= path.join(__dirname,'styles')
        return config
    },
    images: {
      domains: ["source.unsplash.com"],
    },
  }