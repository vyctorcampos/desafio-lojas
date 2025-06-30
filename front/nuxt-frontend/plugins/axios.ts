export default defineNuxtPlugin(() => {
  const api = $fetch.create({
    baseURL: 'http://localhost:8080',
  })

  return {
    provide: {
      api,
    },
  }
})
