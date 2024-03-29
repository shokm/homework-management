require('dotenv').config()
export default {
  // Enable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: true,

  // Target: https://go.nuxtjs.dev/config-target
  target: 'static',

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Studule - 宿題管理アプリ',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/stylelint
    '@nuxtjs/stylelint-module',
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    // https://go.nuxtjs.dev/pwa
    '@nuxtjs/pwa',
    '@nuxtjs/dotenv',
    '@nuxtjs/auth-next',
    '@nuxtjs/dayjs',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // baseURL: '',
  },

  // PWA module configuration: https://go.nuxtjs.dev/pwa
  pwa: {
    manifest: {
      lang: 'ja',
    },
  },

  auth: {
    cookie: true,
    // localStorage: true,
    strategies: {
      local: {
        user: {
          property: false,
        },
        endpoints: {
          login: {
            url: '/v1/auth/login',
            method: 'post',
            propertyName: 'token',
          },
          // logout: { url: '/api/v1/auth/logout', method: 'post' },
          logout: false,
          user: { url: '/v1/auth/user', method: 'get', propertyName: false },
        },
        // tokenRequired: true,
        // tokenType: 'bearer'
      },
    },
  },

  publicRuntimeConfig: {
    API_URL: process.env.API_URL,
    MICROCMS_API_URL: process.env.MICROCMS_API_URL,
    MICROCMS_API_KEY: process.env.MICROCMS_API_KEY,
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
}
