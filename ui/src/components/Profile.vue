<template>
        <div id='profile_component' class='col-xs-12'>
            <h1 class='text-center margin-lg'>{{ profile.username }}'s User Profile</h1>
            <div class='col-xs-12'>Email: {{ profile.email }}</div>
            <div class='col-xs-12'>Firstname: {{ profile.firstName }}</div>
            <div class='col-xs-12'>ID: {{ profile.id }}</div>
            <div class='col-xs-12'>Admin: {{ profile.user_is_admin }}</div>
            <div class='col-xs-12'>Lastname: {{ profile.lastName }}</div>
            <div class='col-xs-12'>Username: {{ profile.username }}</div>
            <hr />
        </div>
</template>

<script>
  import Cookies from "cookie";
  import axios from "axios";

  export default {
    name: 'Profile',
    data: function() {
        return {profile: {}, password: {}};
    },
    beforeCreate: function() {
        var self = this;

        // if(Cookies.get('username') && Cookies.get('user_id')) {
            axios.get('http://api-vf.localhost/user/' + this.$route.params.user, {})
            .then(function (response) {
                self.profile = response.data.result.data.user;
                // self.password = response.data[0].user_password.replace(/./g, "*");
            })
            .catch(function (error) {
                console.log(error);
            });
        // }
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
