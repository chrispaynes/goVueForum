<template>
        <div id='registration_component' method='POST' onsubmit="completeAndRedirect();return false;" class='col-xs-12'>
            <form id='Login' action='data/commands/Login.php' method='POST'>
                <h1 class='text-center margin-lg'>Forum Login</h1>
                <div class='form-group col-xs-12 col-sm-4'>
                    <label class='col-xs-12 col-sm-4 txt-green text-size-sm'>Username</label>
                    <input class='col-xs-12 col-sm-8 margin-bottom-xs pad-vert-xs' type='text' name='uname' placeholder='Username' required>
                </div>
                <div class='form-group col-xs-12 col-sm-4'>
                    <label class='col-xs-12 col-sm-4 txt-green text-size-sm'>Password</label>
                    <input v-model="pwd" class='col-xs-12 col-sm-8 margin-bottom-xs pad-vert-xs' type='password' placeholder='Password' required>
                </div>
                <div class='form-group col-xs-12 col-sm-2'>
                    <input class='hidden' type='text' name='pwd' v-bind:value='SHA1(pwd)' required>
                    <input class='btn btn-success' type='submit' name='login' value='Login'>
                </div>
            </form>
        </div>
</template>

<script>
  export default {
    name: 'Login',
    props: ['posts'],
    data () {
      return { pwd: '' };
    },
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
