'use strict'

const { default: lambdaPlayground } = require('graphql-playground-middleware-lambda')

console.log('APP_GRAPHQL_ENDPOINT', process.env.APP_GRAPHQL_ENDPOINT)

exports.playgroundHandler = lambdaPlayground({
    endpoint: process.env.APP_GRAPHQL_ENDPOINT
})
