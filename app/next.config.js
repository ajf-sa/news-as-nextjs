const path = require('path')
require("dotenv").config()
module.exports = {
    env:{
        API_URL: process.env.API_URL
    },
    future: {
      webpack5: true,
    },
    webpack:config =>{
        config.resolve.alias['components']= path.join(__dirname,'components')
        config.resolve.alias['public']= path.join(__dirname,'public')
        config.resolve.alias['styles']= path.join(__dirname,'styles')
        return config
    }
  }