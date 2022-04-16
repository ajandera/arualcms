<template>
  <div id="app">
    <section class="main">
      <div class="container-fluid">
        <div class="row">
          <div class="col-2 padding-0 d-none d-sm-block sideBg">
            <div class="container">
            <div class="page-wrapper row">
              <nav id="sidebar" class="sidebar-wrapper">
                <div class="sidebar-content">
                  <div class="sidebar-header">
                    <div class="row">
                      <div class="col-3 col-sm-12">
                        <h1>
                          <font-awesome-icon icon="user"/>
                        </h1>
                      </div>
                      <div class="col-9 col-sm-12">
                        <div class="user-info">
                          <span class="user-name" v-if="loggedUser">{{ loggedUser }}</span>
                          <a href="#" @click="logout()">Sign out</a>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="sidebar-menu">
                    <ul>
                      <li class="header-menu">
                        <span>General</span>
                      </li>
                      <li class="sidebar-dropdown" v-for="(item, index) in general" :key="'g'+index">
                        <router-link :to="item.component">{{ item.label }}</router-link>
                      </li>
                      <li class="header-menu">
                        <span>Extra</span>
                      </li>
                      <li v-for="(item, index) in extra" :key="'e'+index">
                        <router-link :to="item.component">{{ item.label }}</router-link>
                      </li>
                    </ul>
                  </div>
                </div>
                <div class="sidebar-footer text-center">
                  <div class="row">
                    <div class="col-12">
                        <p>2021 &copy; v0.9 <a href="https://ajandera.com" target="_blank">ajandera.com</a></p>
                    </div>
                  </div>
                </div>
              </nav>
            </div>
            </div>
          </div>
          <div class="col-xs-12 col-sm-10">
            <nav class="navbar navbar-expand-lg navbar-light bg-light" id="header">
              <a class="navbar-brand" href="/">arualCMS</a>
              <div class="btn-group" role="group" aria-label="Language switch">
                <button
                    v-on:click="setLanguage(lang)"
                    v-bind:class="{'btn btn-default': lang !== language, 'btn btn-primary': lang === language}"
                    v-for="(lang, index) in languages" v-bind:key="index" >{{ lang }}
                </button>
              </div>
              <a href="#" class="icon d-block d-sm-none" v-on:click="hamburger()">
                <font-awesome-icon :icon="['fas', 'bars']"/>
              </a>
              <div class="topnav d-block d-sm-none">
                <div id="mobileMenu" v-if="menu === true">
                  <ul>
                    <li class="sidebar-dropdown" v-for="(item, index) in general" :key="'gm'+index">
                      <router-link :to="item.component">{{ item.label }}</router-link>
                    </li>
                    <li v-for="(item, index) in extra" :key="'em'+index">
                      <router-link :to="item.component">{{ item.label }}</router-link>
                    </li>
                  </ul>
                </div>
              </div>
            </nav>
            <router-view :language="language" :languages="languages" :loggedUser="loggedUser"></router-view>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: 'App',
  data() {
    return {
      general: [
        {
          component: {name: 'posts'},
          label: "Posts"
        },
        {
          component: {name: 'texts'},
          label: "Texts"
        },
        {
          component: {name: 'files'},
          label: "Files"
        },
      ],
      extra: [
        {
          component: {name: 'settings'},
          label: "Settings"
        },
        {
          component: {name: 'languages'},
          label: "Languages"
        },
        {
          component: {name: 'users'},
          label: "Users"
        }
      ],
      loggedUser: window.localStorage.getItem("user"),
      message: null,
      messageClass: null,
      error: null,
      menu: false,
      language: null,
      languages: []
    }
  },
  components: {},
  mounted() {
    if (this.loggedUser === null) {
      this.$router.push({name: 'sign'});
    } else {
      if (this.$router.currentRoute.name !== 'admin/posts') {
        this.$router.push({name: 'posts'});
      }
    }
    this.getDefaultLanguage();
  },
  methods: {
    logout() {
      window.localStorage.removeItem("userId");
      window.localStorage.removeItem("user");
      window.localStorage.removeItem("jwt");
      this.loggedUser = null;
      this.$router.push({name: 'sign'});
    },
    hamburger() {
      this.menu = !this.menu;
    },
    getDefaultLanguage() {
      axios.get(this.$hostname + "languages")
          .then(response => {
            if (response.data.success === true) {
              this.languages = response.data.languages.map(item => item = item.key);
              this.language = response.data.languages.find(item => item.default === true).key;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    setLanguage(lang) {
      this.language = lang;
    }
  },
  watch: {
    $route(to, from) {
      this.menu = false;
    }
  }
}
</script>

<style lang="css" scoped>
@import "https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/css/bootstrap.min.css";

@keyframes swing {
  0% {
    transform: rotate(0deg);
  }
  10% {
    transform: rotate(10deg);
  }
  30% {
    transform: rotate(0deg);
  }
  40% {
    transform: rotate(-10deg);
  }
  50% {
    transform: rotate(0deg);
  }
  60% {
    transform: rotate(5deg);
  }
  70% {
    transform: rotate(0deg);
  }
  80% {
    transform: rotate(-5deg);
  }
  100% {
    transform: rotate(0deg);
  }
}

@keyframes sonar {
  0% {
    transform: scale(0.9);
    opacity: 1;
  }
  100% {
    transform: scale(2);
    opacity: 0;
  }
}

.page-wrapper .sidebar-wrapper,
.sidebar-wrapper .sidebar-dropdown > a:after,
.sidebar-wrapper .sidebar-menu .sidebar-dropdown .sidebar-submenu li a:before,
.sidebar-wrapper ul li a i,
.sidebar-wrapper .sidebar-menu ul li a {
  -webkit-transition: all 0.3s ease;
  -moz-transition: all 0.3s ease;
  -ms-transition: all 0.3s ease;
  -o-transition: all 0.3s ease;
  transition: all 0.3s ease;
}

/*----------------page-wrapper----------------*/
.sideBg {
  background: #1e2229;
}

.page-wrapper {
  height: 100vh;
  background: #1e2229;
  color: #fff;
}

.sidebar-wrapper {
  height: 100%;
  max-height: 100%;
  z-index: 999;
  position: fixed;
}

.sidebar-wrapper ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.sidebar-wrapper a {
  text-decoration: none;
  color: silver;
}

.sidebar-wrapper a:hover {
  text-decoration: underline;
  color: green;
}

.sidebar-wrapper a.router-link-active {
  color: green;
}

/*----------------sidebar-content----------------*/

.sidebar-content {
  max-height: calc(100% - 30px);
  height: calc(100% - 30px);
  overflow-y: auto;
  position: relative;
}

/*--------------------sidebar-header----------------------*/

.sidebar-wrapper .sidebar-header {
  padding: 20px;
  overflow: hidden;
}

.sidebar-wrapper .sidebar-header .user-info {
  float: left;
}

.sidebar-wrapper .sidebar-header .user-info .user-name {
  margin-top: 10px;
  font-size: 12px;
}

.sidebar-wrapper .sidebar-header .user-info > span {
  display: block;
}

/*----------------------sidebar-menu-------------------------*/

.sidebar-wrapper .sidebar-menu {
  padding-bottom: 10px;
}

.sidebar-wrapper .sidebar-menu .header-menu span {
  font-weight: bold;
  font-size: 14px;
  padding: 15px 20px 5px 20px;
  display: inline-block;
}

.sidebar-wrapper .sidebar-menu ul li a {
  display: inline-block;
  width: 100%;
  text-decoration: none;
  position: relative;
  padding: 8px 30px 8px 20px;
}

.sidebar-wrapper .sidebar-menu ul li a i {
  margin-right: 10px;
  font-size: 12px;
  width: 30px;
  height: 30px;
  line-height: 30px;
  text-align: center;
  border-radius: 4px;
}

.sidebar-wrapper .sidebar-menu ul li a:hover > i::before {
  display: inline-block;
  animation: swing ease-in-out 0.5s 1 alternate;
}

.sidebar-wrapper .sidebar-menu .sidebar-dropdown .sidebar-submenu ul {
  padding: 5px 0;
}

.sidebar-wrapper .sidebar-menu .sidebar-dropdown .sidebar-submenu li {
  padding-left: 25px;
  font-size: 13px;
}

.sidebar-wrapper .sidebar-menu .sidebar-dropdown.active > a:after {
  transform: rotate(90deg);
  right: 17px;
}

/*------scroll bar---------------------*/

::-webkit-scrollbar {
  width: 5px;
  height: 7px;
}

::-webkit-scrollbar-button {
  width: 0px;
  height: 0px;
}

::-webkit-scrollbar-thumb {
  background: #525965;
  border: 0px none #ffffff;
  border-radius: 0px;
}

::-webkit-scrollbar-thumb:hover {
  background: #525965;
}

::-webkit-scrollbar-thumb:active {
  background: #525965;
}

::-webkit-scrollbar-track {
  background: transparent;
  border: 0px none #ffffff;
  border-radius: 50px;
}

::-webkit-scrollbar-track:hover {
  background: transparent;
}

::-webkit-scrollbar-track:active {
  background: transparent;
}

::-webkit-scrollbar-corner {
  background: transparent;
}

.modal-content {
  border: none;
}

.padding-0 {
  padding: 0;
}

/* Style the navigation menu */
.topnav {
  overflow: hidden;
  background-color: #333;
  width: 200px;
  z-index: 999;
}

/* Style navigation menu links */
.topnav a {
  color: white;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
  display: block;
}

/* Add a grey background color on mouse-over */
.topnav a:hover {
  background-color: #ddd;
  color: black;
}

#mobileMenu {
  position: absolute;
  background: #000;
  width: 216px;
  top: 53px;
  right: 0;
}

#mobileMenu ul {
  list-style: none;
  padding-left: 0;
}
</style>
