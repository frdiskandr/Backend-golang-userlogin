<template>
    <div class="card-container">
        <h1>Register</h1>
        <form @submit.prevent="Submit">
            <label for="username">Username</label>
            <input type="text" name="username" id="username"
            v-model="username" placeholder="Username">
            <label for="email">Email</label>
            <input type="text" name="email" id="username"
            v-model="email" placeholder="Email">
            <label for="password">Password</label>
            <input type="password" name="password" id="password"
            v-model="password" placeholder="Password">
            <input type="submit" value="Sign in" id="submit" 
            :disabled="!username || !email || !password">
            <p>have an account? <router-link to="/">Sign in</router-link></p>
        </form>
    </div>
</template>

<script>
export default {
    name: 'registerView',

    data() {
        return {
            userRegister: {
                username: '',
                email: '',
                password: ''
            }
        }
    },

    methods: {
        Submit() {
            this.userRegister = { username: this.username, email: this.email, password: this.password }
            console.log(this.userRegister)
        }
    }
}
</script>

<style scoped>
.card-container{
    padding: 1rem;
    margin: auto;
    display: grid;
    background-color: rgb(31, 30, 30);
    color: aliceblue;
    border-radius: 10px;
    overflow: hidden;
    min-width: 300px;
    max-width: 500px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
h1{
    text-align: center;
}
form{
    display:grid;
    grid-template-columns: 1fr;
}
label{
    padding-left: 20px;
}
input{
    margin: 1em;
    padding: 1em;
    border: none;
    border-radius: 5px;
}
#submit{
    margin: 1em;
    padding: 1em;
    border: none;
    border-radius: 5px;
    background-color: rgb(60, 255, 0);
    color: rgb(18, 18, 19);
    transition: 0.2s ease-in-out;
    cursor: pointer;
    }
#submit:hover{
     background-color: rgb(33, 139, 1);
     transform: scale(0.99);
}
</style>