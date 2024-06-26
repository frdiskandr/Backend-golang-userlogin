import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import axios from 'axios'


axios.defaults.baseURL = 'http://localhost:8080'
axios.defaults.headers.common['Authorization'] = `Bearer ${localStorage.getItem('token')}`
const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
