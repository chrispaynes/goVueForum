<template>
  <div id="app">
    <div class='pad-bottom-md'>
        <Header></Header>
        <router-view></router-view>
        <Footer></Footer>
    </div>
  </div>
</template>

<script>
  import Footer from './components/Footer';
  import Header from './components/Header';

  export default {
    name: 'app',
    components: {
      Header,
      Footer
    }
  }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

* {
  text-decoration: none !important;
  list-style-type: none;
}

.bdr {
  border: solid;
}


/* SUBFORUMS COMPONENT */

.sf_header {
  font-weight: bold;
}

.thread_headline {
  padding: 0 200px;
  line-height: 125%;
  font-size: 48px;
}

.thread_byline {
  line-height: 125%;
  font-size: 40px;
  font-family: serif;
}

.reply_label {
  font-family: serif;
  font-size: 24px;
}

#post_listing li:first-child {
  border-bottom: 3px solid #34495e;
  border-top: 3px solid #34495e;
}


/* HEADER COMPONENT */

#nav_component a {
  text-align: center;
  margin-top: 48px;
}

#footer_component > div > a:hover {
  font-weight: bold;
  color: #34495e;
}

@media (max-width: 576px) {
  #logo {
    text-align: center;
  }
}

@media (max-width: 767px) {
  body {
    line-height: 150%;
  }
  /* INDEX COMPONENT */
  #index_label-container {
    display: none;
  }
}

@media (max-width: 992px) {
  #nav_component a {
    margin-bottom: 12px;
    margin-top: 24px;
  }
  .post_author,
  .post_date {
    text-align: center;
  }
  .post_date {
    margin-bottom: 32px;
  }
  /* THREAD COMPONENT */
  .thread_headline {
    padding: 0 20px;
  }
  #new_thread_component form div input,
  #new_thread_component form div textarea {
    margin-bottom: 32px;
  }
  #new_thread_component form div label {
    text-align: center;
  }
}


/* Utility Styles */

.pad-0 {
  padding: 0;
}

.pad-xs {
  padding: 12px;
}

.pad-sm {
  padding: 16px;
}

.pad-md {
  padding: 32px;
}

.pad-lg {
  padding: 48px;
}

.pad-top-xs,
.pad-vert-xs {
  padding-top: 8px;
}

.pad-top-sm,
.pad-vert-sm {
  padding-top: 16px;
}

.pad-top-md,
.pad-vert-md {
  padding-top: 32px;
}

.pad-top-lg,
.pad-vert-lg {
  padding-top: 48px;
}

.pad-bottom-xs,
.pad-vert-xs {
  padding-bottom: 8px;
}

.pad-bottom-sm,
.pad-vert-sm {
  padding-bottom: 16px;
}

.pad-bottom-md,
.pad-vert-md {
  padding-bottom: 32px;
}

.pad-bottom-lg,
.pad-vert-lg {
  padding-bottom: 48px;
}

.pad-left-sm,
.pad-horiz-sm {
  padding-left: 16px;
}

.pad-left-md,
.pad-horiz-md {
  padding-left: 32px;
}

.pad-left-lg,
.pad-horiz-lg {
  padding-left: 48px;
}

.pad-right-sm,
.pad-horiz-sm {
  padding-right: 16px;
}

.pad-right-md,
.pad-horiz-md {
  padding-right: 32px;
}

.pad-right-lg,
.pad-horiz-lg {
  padding-right: 48px;
}

.margin-0 {
  margin: 0;
}

.margin-sm {
  margin: 16px;
}

.margin-md {
  margin: 32px;
}

.margin-lg {
  margin: 48px;
}

.margin-bottom-0 {
  margin-bottom: 0;
}

.margin-bottom-xs,
.margin-vert-xs {
  margin-bottom: 12px;
}

.margin-bottom-sm,
.margin-vert-sm {
  margin-bottom: 16px;
}

.margin-bottom-md,
.margin-vert-md {
  margin-bottom: 32px;
}

.margin-bottom-lg,
.margin-vert-lg {
  margin-bottom: 48px;
}

.margin-bottom-xl {
  margin-bottom: 128px;
}

.margin-top-xs,
.margin-vert-xs {
  margin-top: 12px;
}

.margin-top-sm,
.margin-vert-sm {
  margin-top: 16px;
}

.margin-top-md,
.margin-vert-md {
  margin-top: 32px;
}

.margin-top-lg,
.margin-vert-lg {
  margin-top: 48px;
}

.margin-top-xl {
  margin-top: 128px;
}

.margin-left-xs,
.margin-horiz-xs {
  margin-left: 8px;
}

.margin-left-sm,
.margin-horiz-sm {
  margin-left: 16px;
}

.margin-left-md,
.margin-horiz-md {
  margin-left: 32px;
}

.margin-left-lg,
.margin-horiz-lg {
  margin-left: 48px;
}

.margin-right-xs,
.margin-horiz-xs {
  margin-right: 8px;
}

.margin-right-sm,
.margin-horiz-sm {
  margin-right: 16px;
}

.margin-right-md,
.margin-horiz-md {
  margin-right: 32px;
}

.margin-right-lg,
.margin-horiz-lg {
  margin-right: 48px;
}

a.postlink {
  font-size: 16px;
  text-decoration: none;
}

.txt-green,
a:hover {
  color: #42b983;
}

.txt-blue {
  color: #34495e;
}

.txt-gray {
  color: #7f8c8d;
}

.text-size-xs {
  font-size: 14px;
}

.text-size-sm {
  font-size: 16px;
}

.text-size-md {
  font-size: 32px;
}

.text-size-lg {
  font-size: 48px;
}

.bg-green {
  background-color: #42b983;
}

.bg-blue {
  background-color: #34495e;
}

.bg-gray {
  background-color: #e9e9e9;
}

.bold {
  font-weight: bold;
}

.inline {
  display: inline;
}

hr {
  border-top: 2px solid #e9e9e9;
}

.hr-bdr-4 {
  border-top: 4px solid #e9e9e9;
}

.hr-bdr-8 {
  border-top: 8px solid #e9e9e9;
}

</style>
