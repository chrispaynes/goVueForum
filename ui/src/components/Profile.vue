<template>
        <div id='profile_component' class='col-xs-12'>
            <h1 class='text-center margin-lg'>{{ profile.user_username }}'s User Profile</h1>
            <div class='col-xs-12'>Email: {{ profile.user_email }}</div>
            <div class='col-xs-12'>Firstname: {{ profile.user_firstname }}</div>
            <div class='col-xs-12'>ID: {{ profile.user_id_PK }}</div>
            <div class='col-xs-12'>Admin: {{ profile.user_is_admin }}</div>
            <div class='col-xs-12'>Lastname: {{ profile.user_lastname }}</div>
            <div class='col-xs-12'>Password: {{ password }}</div>
            <div class='col-xs-12'>Username: {{ profile.user_username }}</div>
            <hr />
        </div>
</template>

<script>
  import Cookies from "cookie";

  export default {
    name: 'Profile',
    data: function() {
        return {profile: this.profile, password: self.password};
    },
    beforeCreate: function() {
        var self = this;

        if(Cookies.get('username') && Cookies.get('user_id')) {
            axios.get('data/queries/User.php', {
                params: {
                    user: Cookies.get('username') || Cookies.get('user_id'),
                },
            })
            .then(function (response) {
                self.profile = response.data[0];
                self.password = response.data[0].user_password.replace(/./g, "*");
                // pad short passwords
                while(self.password.length < 10) {
                    self.password += '*';
                }
            })
            .catch(function (error) {
                console.log(error);
            });
        }
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
