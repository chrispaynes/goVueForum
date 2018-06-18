<template>
        <div id='reply_component' class="col-xs-12 text-center pad-vert-md">
            <form action='http://api-vf.localhost/reply' method="POST">
                <div class="form-group col-xs-12 margin-bottom-md">
                    <label class="col-sm-12 margin-bottom-sm txt-green reply_label">Reply</label>
                    <textarea rows="12" class="col-xs-12" type="text" name="body" placeholder="" required />
                    <input class='hidden' type='text' name='thread_id' v-bind:value='this.thread_id' required>
                    <input class='hidden' type='text' name='username' v-bind:value='this.username' required>
                    <input class='hidden' type='text' name='user_id' v-bind:value='this.user_id' required>
                    <input class='hidden' type='text' name='timestamp' v-bind:value='new Date()' required>
                </div>
                <div class='col-xs-12'>
                    <input class='btn btn-primary margin-horiz-xs' type=submit name='new_reply' value='Reply'>
                    <router-link class='btn btn-success margin-horiz-xs' to='/thread/new'>Start New Thread</router-link>
                </div>
            </form>
            <hr />
        </div>
</template>

<script>
  import Cookies from "cookie";

  export default {
    name: 'Reply',
    data: function() {
        return { username: this.username, user_id: this.user_id, thread_id: this.thread_id };
    },
    beforeCreate: function() {
        var self = this;

        self.thread_id = this.$route.query.thread;
        if(Cookies.length >= 2 && Cookies.get('user_id') && Cookies.get('username')) {
            this.username = Cookies.get('username').length > 0  ? Cookies.get('username') : false;
            this.user_id = Cookies.get('user_id').length > 0 ? Cookies.get('user_id') : false;
            this.isLoggedOut = false;
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
