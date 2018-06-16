<template>
        <div class='row'>
            <ul class='index_component col-xs-12 container-fluid'>
                <li class='col-xs-12'>
                    <div class='row txt-green bold'>
                        <div class='col-xs-offset-1 col-xs-2 col-md-5'>Topic</div>
                        <div class='col-xs-2 col-md-2'>Author</div>
                        <div class='col-xs-4 col-md-4'>Last Updated</div>
                    </div>
                    <hr class='hr-bdr-4 margin-top-xs'/>
                </li>
                <li v-for='(p, index) in posts'
                    v-bind:subforum='p'
                    v-bind:index='index'
                    v-bind:key='p.id'
                    class='col-xs-12'
                >
                    <div class='row pad-vert-sm'>
                        <div class='col-xs-1 col-md-1'>#{{ index + 1 }}.</div>
                        <div class='col-xs-11 col-md-5'>
                            <router-link :to="{ path: '/thread/', query: { thread: p.post_thread_id }}" class='postlink'><strong>{{ p.thread_name }}</strong></router-link>
                        </div>
                        <div class='col-xs-offset-1 col-xs-12 col-sm-3 col-md-offset-0 col-md-2'>by {{ p.author_username }}</div>
                        <div class='col-xs-offset-1 col-xs-12 col-sm-8 col-md-offset-0 col-md-4'>on {{ p.post_last_updated }}</div>
                    </div>
                    <hr v-if="index != (posts.length - 1)"/>
                </li>
            </ul>
        </div>
</template>

<script>
  import axios from "axios";

  export default {
    name: 'Index',
    props: ['posts'],
    data: function() {
        return { posts: [], most_recent_post: [] };
    },
    beforeCreate: function() {
        var self = this;

        axios.get('./data/queries/Post.php', {})
        .then(function(response) {
            self.posts = response.data.filter(function(p) {return p.post_is_reply == 'false';});
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
