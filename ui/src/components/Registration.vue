<template>
    <div id='registration_component' class='col-xs-12'>
        <form id='registration_form' action='http://api-vf.localhost/register' method='POST'>
            <h1 class='text-center margin-lg'>Forum Registration</h1>
            <div class='form-group col-xs-12 col-sm-6'>
                <label class='col-xs-12 col-sm-3 txt-green text-size-sm'>Username</label>
                <input class='col-xs-12 col-sm-6 margin-bottom-xs pad-vert-xs' type='text' name='uname' placeholder='Username' required>
            </div>
            <div class='form-group col-xs-12 col-sm-6'>
                <label class='col-xs-12 col-sm-3 txt-green text-size-sm'>Password</label>
                <input class='col-xs-12 col-sm-6 margin-bottom-xs pad-vert-xs' type='text' name='pwd' placeholder='Password' required>
            </div>
            <div class='form-group col-xs-12 col-sm-6'>
                <label class='col-xs-12 col-sm-3 txt-green text-size-sm'>First Name</label>
                <input class='col-xs-12 col-sm-6 margin-bottom-xs pad-vert-xs' type='text' name='fname' placeholder='First Name' required>
            </div>
            <div class='form-group col-xs-12 col-sm-6'>
                <label class='col-xs-12 col-sm-3 txt-green text-size-sm'>Last Name</label>
                <input class='col-xs-12 col-sm-6 margin-bottom-xs pad-vert-xs' type='text' name='lname' placeholder='Last Name' required>
            </div>
            <div class='form-group col-xs-12 col-sm-6'>
                <label class='col-xs-12 col-sm-3 txt-green text-size-sm'>Email</label>
                <input class='col-xs-12 col-sm-6 margin-bottom-xs pad-vert-xs' type='email' name='email' placeholder='Email' required>
            </div>
            <div class='form-group col-xs-12 col-sm-12 text-center'>
                <input class='btn btn-success' type='submit' name='register' value='Register'>
            </div>
        </form>
    </div>
</template>

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
