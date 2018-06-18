<template>
        <div id='registration_component' class='col-xs-12'>
            <form id='Login' @submit.prevent="onSubmit">
                <h1 class='text-center margin-lg'>Forum Login</h1>
                <div class='form-group col-xs-12 col-sm-4'>
                    <label class='col-xs-12 col-sm-4 txt-green text-size-sm'>Username</label>
                    <input v-model="login.username" class='col-xs-12 col-sm-8 margin-bottom-xs pad-vert-xs' type='text' name='uname' placeholder='Username' required>
                </div>
                <div class='form-group col-xs-12 col-sm-4'>
                    <label class='col-xs-12 col-sm-4 txt-green text-size-sm'>Password</label>
                    <input v-model="login.password" class='col-xs-12 col-sm-8 margin-bottom-xs pad-vert-xs' type='password' placeholder='Password' required>
                </div>
                <div class='form-group col-xs-12 col-sm-2'>
                    <input class='btn btn-success' type='submit' name='login' value='Login'>
                </div>
            </form>
        </div>
</template>

<script>
  export default {
    name: 'Login',
    data () {
      return { login: {username: '', password: ''} };
    },
    methods: {
        onSubmit () {
          console.log("THIS: ", this);
            this.$http.post('http://api-vf.localhost/login', this.login);
        }
    }
  }
</script>


<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
