<template>
  <div class="card-container">
    <h1>Sign in</h1>
    <form @submit.prevent="login">
      <label for="email">Email</label>
      <input type="email" name="email" id="email" v-model="email" placeholder="Email" />
      <label for="password">Password</label>
      <input
        type="password"
        name="password"
        id="password"
        v-model="password"
        placeholder="Password"
      />
      <button type="submit" value="Sign in" id="submit">Sign in</button>
      <p>Don't have an account? <router-link to="/register">Register</router-link></p>
    </form>
    <div v-if="error">{{ error }}</div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'loginView',
  data() {
    return {
      userLogin: {
        email: '',
        password: '',
        error: ''
      }
    }
  },

  methods: {
    async login() {
          try {
        console.log("loding")
        const response = await axios.post('/login', {
          email: this.email,
          password: this.password
        })
        localStorage.setItem('token', response.data.token)
        axios.defaults.headers.common['Authorization'] = `Bearer ${localStorage.getItem('token')}`
        this.$router.push('/dashboard')
      } catch (error) {
        this.error = 'login faild : ' + error.response.data.message
      }
    }
  }
}
</script>

<style scoped>
.card-container {
  padding: 1rem;
  margin: auto;
  display: grid;
  background-color: rgb(31, 30, 30);
  color: aliceblue;
  border-radius: 10px;
  overflow: hidden;
  min-width: 300px;
  max-width: 500px;
  box-shadow:
    0 4px 8px 0 rgba(0, 0, 0, 0.2),
    0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
h1 {
  text-align: center;
}
form {
  display: grid;
  grid-template-columns: 1fr;
}
label {
  padding-left: 20px;
}
input {
  margin: 1em;
  padding: 1em;
  border: none;
  border-radius: 5px;
}
#submit {
  margin: 1em;
  padding: 1em;
  border: none;
  border-radius: 5px;
  background-color: rgb(60, 255, 0);
  color: rgb(18, 18, 19);
  transition: 0.2s ease-in-out;
  cursor: pointer;
}
#submit:hover {
  background-color: rgb(33, 139, 1);
  transform: scale(0.99);
}
</style>
