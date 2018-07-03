<template>
        <div class='row'>
            <ul class='index_component col-xs-12 container-fluid'>
                <li class='col-xs-12'>
                    <div class='row txt-green bold'>
                        <div class='col-xs-1 col-sm-1'>ID</div>
                        <div class='col-xs-3 col-sm-2'>Username</div>
                        <div class='col-xs-3 col-sm-1'>Post Count</div>
                        <div class='col-xs-3 col-sm-2'>First Name</div>
                        <div class='col-xs-4 col-sm-1'>Last Initial</div>
                        <div class='col-xs-11 col-xs-offset-1 col-sm-offset-0 col-sm-3'>Email</div>
                        <div class='col-xs-3 col-sm-2'>Location</div>
                        <!-- <div class='col-xs-11 col-xs-offset-1 col-sm-offset-0 col-sm-3'>Access Level</div> -->
                    </div>
                    <hr class='hr-bdr-4 margin-top-xs'/>
                </li>
                <li v-for='(u, index) in users'
                    v-bind:key='u.id'
                    v-bind:index='index'
                    class='col-xs-12'
                >
                    <div class='row pad-vert-sm'>
                        <div class='col-xs-1 col-sm-1'>#{{ u.id }}</div>
                        <div class='col-xs-2 col-sm-2 text-size-sm'>
                          <router-link :to="{ path: `/profile/${u.id}`}">
                            <strong>{{ u.username }}</strong>
                          </router-link>
                        </div>
                        <div class='col-xs-3 col-sm-1'>{{ u.postCount }}</div>
                        <div class='col-xs-3 col-sm-2'>{{ u.firstName }}</div>
                        <div class='col-xs-2 col-sm-1'>{{ u.lastName[0] }}.</div>
                        <div class='col-xs-offset-1 col-xs-4 col-sm-offset-0 col-sm-3'>
                          {{ u.email }}
                        </div>
                        <div class='col-xs-3 col-sm-2'>{{ u.location }}</div>
                        <!-- <div class='col-xs-offset-1 col-xs-12 col-sm-offset-0 col-sm-3 txt-green' v-if='u.user_is_admin === "true"'><strong><i>Admin</i></strong></div>
                        <div class='col-xs-offset-1 col-xs-12 col-sm-offset-0 col-sm-3 txt-blue' v-if='u.user_is_admin === "false"'><i>User</i></div> -->
                    </div>
                    <hr />
                </li>
            </ul>
        </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'UserIndex',
    props: ['users'],
    data: function() {
        return { };
    },
    beforeCreate: function() {
        var self = this;

        axios.get('http://api-vf.localhost/users', {})
        .then(function(response) {
            self.users = response.data.result.data.users;
        })
        .catch(function(error) {
            console.log(error);
        });
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
