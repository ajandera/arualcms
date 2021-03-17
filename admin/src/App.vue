<template>
  <div id="app">
    <section class="main">
      <div class="container-fluid">
        <div class="row">
          <div class="col-2 padding-0 d-none d-sm-block">
            <div class="page-wrapper" v-if="loggedUser">
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
                      <a href="#" @click="logout()">Sign out</a>
                    </div>
                  </div>
                  <!-- sidebar-header  -->
                  <div class="sidebar-menu">
                    <ul>
                      <li class="header-menu">
                        <span>General</span>
                      </li>
                      <li class="sidebar-dropdown" v-for="item in general" :key="item.component">
                        <router-link :to="item.component">{{ item.label }}</router-link>
                      </li>
                      <li class="header-menu">
                        <span>Extra</span>
                      </li>
                      <li v-for="item in extra" :key="item.component">
                        <router-link :to="item.component">{{ item.label }}</router-link>
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
          <div class="col-xs-10 col-sm-10">
            <nav class="navbar navbar-expand-lg navbar-light bg-light" id="header" v-if="loggedUser">
              <a class="navbar-brand" href="/">arualCMS</a>
              <div class="topnav d-block d-sm-none">
                <!-- Navigation links (hidden by default) -->
                <div id="myLinks" v-if="menu === true">
                  <ul>
                    <li class="sidebar-dropdown" v-for="item in general" :key="item.component">
                      <router-link :to="item.component">{{ item.label }}</router-link>
                    </li>
                    <li v-for="item in extra" :key="item.component">
                      <router-link :to="item.component">{{ item.label }}</router-link>
                    </li>
                  </ul>
                </div>
                <!-- "Hamburger menu" / "Bar icon" to toggle the navigation links -->
                <a href="#" class="icon" v-on:click="hamburger()">
                  <font-awesome-icon :icon="['fas', 'bars']" />
                </a>
              </div>
            </nav>
            <router-view></router-view>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      general: [
        {
          component: "/posts",
          label: "Posts"
        },
        {
          component: "/texts",
          label: "Texts"
        },
        {
          component: "/files",
          label: "Files"
        },
      ],
      extra: [
        {
          component: "/settings",
          label: "Settings"
        },
        {
          component: "/users",
          label: "Users"
        }
      ],
      loggedUser: window.localStorage.getItem("user"),
      message: null,
      messageClass: null,
      error: null,
      menu: false
    }
  },
  components: {},
  mounted() {
    if (this.loggedUser === null) {
      this.$router.push('signin');
    } else {
      this.$router.push('posts');
    }
  },
  methods: {
    logout() {
      window.localStorage.removeItem("userId");
      window.localStorage.removeItem("user");
      window.localStorage.removeItem("jwt");
      this.loggedUser = null;
      this.$router.push('signin');
    },
    hamburger() {
      this.menu = !this.menu;
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

/* Style the hamburger menu */
.topnav a.icon {
  background: black;
  display: block;
  position: absolute;
  right: 0;
  top: 0;
  cursor: pointer;
  z-index: 999;
}

/* Add a grey background color on mouse-over */
.topnav a:hover {
  background-color: #ddd;
  color: black;
}
#myLinks {
  position: absolute;
  background: #000;
  width: 216px;
  top: 53px;
}
#myLinks ul {
  list-style: none;
  padding-left: 0;
}
</style>
