<template>
    <div id='thread_component' class='col-xs-12'>
        <div class='pad-sm text-center'>
            <h1 id='thread_headline' class='txt-green thread_headline'>{{ title }} </h1>
            <h2 class='txt-blue thread_byline'><small>by</small> {{ author }}</h2>
        </div>
        <hr />
        <ul id="post_listing" class='pad-md'>
            <li v-for='(p, index) in posts'
                v-bind:p='p'
                v-bind:index='index'
                v-bind:key='p.id'
                v-bind:class="[index === 0 ? 'bg-gray' : 'col-xs-12', 'col-xs-12 margin-bottom-md post']"
            >
                <div class='fluid-container'>
                    <div class='row' v-if='index == 0'>
                        <div class='col-xs-12 col-md-4'>
                            <div class='post_author'>
                              <router-link :to="{ path: `/profile/${p.author.id}`}">
                                <h3>{{ p.author.username }}</h3>
                              </router-link>
                            </div>
                            <div class='txt-gray post_date'>{{ p.lastUpdatedAt}}</div>
                        </div>
                        <div class='col-xs-12 col-md-8 pad-horiz-md pad-bottom-sm'>{{ p.body }}</div>
                    </div>
                </div>
                <div class='row col-xs-12' v-if='index > 0'>
                    <div class='col-xs-12 col-md-4 txt-blue'>
                        <div class='post_author'>
                          <router-link :to="{ path: `/profile/${p.author.id}`}">
                            <h3>{{ p.author.username }}</h3>
                          </router-link>
                        </div>
                        <div class='txt-gray post_date'>{{ p.lastUpdatedAt }}</div>
                    </div>
                    <div class='col-xs-12 col-md-8 pad-horiz-md'>{{ p.body }}</div>
                    <hr class='col-xs-12'/>
                </div>
            </li>
            <div v-if="isLoggedIn">
                <reply-form></reply-form>
            </div>
            <div v-if="!isLoggedIn">
                <router-link to='/register'>Login</router-link> or
                <router-link to='/register'>Register</router-link> to reply
            </div>
        </ul>
    </div>
</template>

<script>
  import Cookies from "cookie";
  import axios from "axios";
  import moment from "moment";

  export default {
    name: 'Thread',
    data: function() {
        return { posts: [],
            title: '',
            author: '',
            username: this.username,
            user_id: this.user_id,
            isLoggedIn: this.isLoggedIn || false,
        };
    },
    beforeCreate: function() {
        var self = this;

        axios.get('http://api-vf.localhost/thread/' + this.$route.params.thread, {
        })
        .then(function(response) {
            self.posts = response.data.result.data.thread.posts;
            self.title = response.data.result.data.thread.title;
            self.author = response.data.result.data.thread.author.username;

            self.posts.forEach(p => {
              p.lastUpdatedAt = moment(p.lastUpdatedAt).format('MMMM Do, YYYY, h:mm:ss a');
          });
        })
        .catch(function(error) {
            console.log(error);
        });

        if(Cookies.length >= 2 && Cookies.get('user_id') && Cookies.get('username')) {
            this.username = Cookies.get('username').length > 0  ? Cookies.get('username') : false;
            this.user_id = Cookies.get('user_id').length > 0 ? Cookies.get('user_id') : false;
            this.isLoggedIn = true;
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
