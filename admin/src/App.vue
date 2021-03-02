<template>
  <div>
    <div id="app" v-if="loggedUser">
    <section class="main">
      <div class="container-fluid">
        <div class="row">
          <div class="col-2 padding-0">
            <div class="page-wrapper">
              <nav id="sidebar" class="sidebar-wrapper">
                <div class="sidebar-content">
                  <div class="sidebar-header">
                    <div class="user-pic">
                      <img class="img-responsive img-rounded" src="https://raw.githubusercontent.com/azouaoui-med/pro-sidebar-template/gh-pages/src/img/user.jpg"
                           alt="User picture">
                    </div>
                    <div class="user-info">
                      <span class="user-name">{{ loggedUser }}</span>
                      <span class="user-role">Administrator</span>
                      <a href="#" @click="logout">Sign out</a>
                    </div>
                  </div>
                  <!-- sidebar-header  -->
                  <div class="sidebar-menu">
                    <ul>
                      <li class="header-menu">
                        <span>General</span>
                      </li>
                      <li class="sidebar-dropdown" v-for="item in general" :key="item.component">
                        <a href="#" v-on:click="chooseComponent(item.component)" :class="item.component === selected ? 'active' : ''">
                          <span>{{ item.label}}</span>
                        </a>
                      </li>
                      <li class="header-menu">
                        <span>Extra</span>
                      </li>
                      <li v-for="item in extra" :key="item.component">
                        <a href="#" v-on:click="chooseComponent(item.component)" :class="item.component === selected ? 'active' : ''">
                          <span>{{ item.label }}</span>
                        </a>
                      </li>
                    </ul>
                  </div>
                  <!-- sidebar-menu  -->
                </div>
                <!-- sidebar-content  -->
                <div class="sidebar-footer text-center">
                    <p>2021 &copy; <a href="https://ajandera.com" target="_blank">ajandera.com</a></p>
                </div>
              </nav>
            </div>
          </div>
          <div class="col-10">
            <nav class="navbar navbar-expand-lg navbar-light bg-light" id="header">
              <a class="navbar-brand" href="/">arualCMS</a>
            </nav>
            <component :is="selected" class="tab"></component>
          </div>
        </div>
      </div>
    </section>
  </div>
    <div id="sign" v-if="!loggedUser">
      <modal name="sign" height="auto">
        <div class="modal-dialog">
          <!-- Modal content-->
          <div class="modal-content">
            <div class="modal-header">
              <h3 class="modal-title">arualCMS</h3>
            </div>
            <div class="modal-body">
              <div class="row">
                <div class="col-12">
                  <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
                </div>
              </div>
              <div class="row">
                <div class="col-12">
                  <label for="username">Username</label>
                  <input type="text" v-model="username" class="form-control" id="username">
                  <label for="username">Password</label>
                  <input type="password" v-model="password" class="form-control" id="password">
                  <button class="float-right btn btn-success mt-3" v-on:click="login">Sign In</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </modal>
    </div>
  </div>
</template>

<script>

import Settings from '@/components/Settings';
import Posts from "@/components/Posts";
import Texts from "@/components/Texts";
import Users from "@/components/Users";
import axios from "axios";

export default {
  name: 'App',
  data() {
    return {
      selected: "Posts",
      general: [
        {
          component: "Posts",
          label: "Posts"
        },
        {
          component: "Texts",
          label: "Texts"
        }
      ],
      extra: [
        {
          component: "Settings",
          label: "Settings"
        },
        {
          component: "Users",
          label: "Users"
        }
      ],
      loggedUser: window.localStorage.getItem("user"),
      message: null,
      messageClass: null,
      username: null,
      password: null,
      error: null,
    }
  },
  components: {
    Settings,
    Posts,
    Texts,
    Users
  },
  mounted() {
    if (!this.loggedUser) {
      this.show();
    }
  },
  methods: {
    chooseComponent(component) {
      this.selected = component;
    },
    login() {
      axios.post(this.$hostname+"auth", {"username": this.username, "password": this.password})
          .then(response => {
            let res = response.data;
            this.message = res.message;
            if (res.success) {
              this.message = res.message;
              this.messageClass = "alert alert-success";
              this.username = null;
              this.password = null;
              window.localStorage.setItem('userId', res.id);
              this.loggedUser = res.username;
              window.localStorage.setItem("user", res.username)
              this.hide();
            } else {
              this.message = res.message;
              this.messageClass = "alert alert-danger";
            }
          });
    },
    logout() {
      window.localStorage.removeItem("userId");
      window.localStorage.removeItem("user");
      this.loggedUser = false;
      window.location.reload();
    },
    show() {
      this.$modal.show('sign')
    },
    hide() {
      this.$modal.hide('sign')
    },
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

.page-wrapper {
  height: 100vh;
  background: #1e2229;
  color: #fff;
}

.sidebar-wrapper {
  height: 100%;
  max-height: 100%;
  z-index: 999;
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

.sidebar-wrapper a.active {
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

.sidebar-wrapper .sidebar-header .user-pic {
  float: left;
  width: 60px;
  padding: 2px;
  border-radius: 12px;
  margin-right: 15px;
  overflow: hidden;
}

.sidebar-wrapper .sidebar-header .user-pic img {
  object-fit: cover;
  height: 100%;
  width: 100%;
}

.sidebar-wrapper .sidebar-header .user-info {
  float: left;
}

.sidebar-wrapper .sidebar-header .user-info > span {
  display: block;
}

.sidebar-wrapper .sidebar-header .user-info .user-role {
  font-size: 12px;
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
</style>
