<template>
        <div id='new_thread_component' class='col-xs-12'>
            <form action='http://api-vf.localhost/thread/new' method='POST'>
            <h1 class='text-center margin-bottom-lg'>Start A New Thread</h1>
                <div class='form-group col-sm-12'>
                    <label class='col-xs-12 col-sm-2 txt-green'>Thread Title</label>
                    <input class='col-xs-12 col-sm-10' type='text' name='title' placeholder='' required>
                </div>
                <div class='form-group col-sm-12'>
                    <label class='col-xs-12 col-sm-2 txt-green'>Thread Message</label>
                    <textarea rows='12' class='col-xs-12 col-sm-10' type='text' name='body' placeholder='' required></textarea>
                </div>
                <input class='hidden' type='text' name='username' v-bind:value='this.username' required />
                <input class='hidden' type='text' name='user_id' v-bind:value='this.user_id' required>
                <input class='hidden' type='text' name='timestamp' v-bind:value='new Date()' required>
                <div class='form-group col-sm-12 text-center'>
                    <input class='btn btn-success' type='submit' name='new_thread' value='Submit Thread'>
                </div>
            </form>
        </div>
</template>

<script>
  import Cookies from "cookie";

  export default {
    name: 'NewThread',
    props: ['posts'],
    data: function() {
        return { username: this.username, user_id: this.user_id };
    },
    beforeCreate: function() {
        var self = this;

        if(Cookies.length > 0 && Cookies.get('username').length > 0 && Cookies.get('user_id').length > 0) {
            self.username = Cookies.get('username');
            self.user_id = Cookies.get('user_id');
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
