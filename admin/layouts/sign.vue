<template>
  <v-app light>
    <!--====== HEADER PART START ======-->
    <header class="header-area">
      <div class="navbar-area">
        <div class="container">
          <div class="row">
            <div class="col-lg-12">
              <nav class="navbar navbar-expand-lg">
                <a class="navbar-brand" href="/">
                  <img src="/logosm.png" alt="arualCMS" class="logo"> {{  title }}
                </a>
                <button
                  class="navbar-toggler"
                  type="button"
                  data-bs-toggle="collapse"
                  data-bs-target="#navbarSupportedContent"
                  aria-controls="navbarSupportedContent"
                  aria-expanded="false"
                  aria-label="Toggle navigation"
                >
                  <span class="toggler-icon"> </span>
                  <span class="toggler-icon"> </span>
                  <span class="toggler-icon"> </span>
                </button>

                <div
                  class="collapse navbar-collapse sub-menu-bar"
                  id="navbarSupportedContent"
                >
                  <ul id="nav" class="navbar-nav ms-auto">
                    <li class="nav-item">
                      <a class="page-scroll active" href="#home">Home</a>
                    </li>
                    <li class="nav-item">
                      <a class="page-scroll" href="#features">Features</a>
                    </li>
                    <li class="nav-item">
                      <a class="page-scroll" href="#about">About</a>
                    </li>
                    <li class="nav-item">
                      <a class="page-scroll" href="#facts">Why</a>
                    </li>
                  </ul>
                </div>
                <!-- navbar collapse -->

                <div class="navbar-btn d-none d-sm-inline-block">
                  <v-dialog
                    v-model="dialogReg"
                    width="500"
                  >
                    <template v-slot:activator="{ on, attrs }">
                      <v-btn
                        text
                        v-bind="attrs"
                        v-on="on"
                        class="main-btn">
                        Sign Up
                      </v-btn>
                    </template>
                    <v-card>
                      <v-card-title class="text-h5 signup white--text lighten-2">
                        Sign Up
                      </v-card-title>
                      <v-card-text>
                        <v-form>
                          <v-text-field
                            prepend-icon="person"
                            name="name"
                            label="Name"
                            type="text"
                            v-model="user.name"
                          ></v-text-field>
                          <v-text-field
                            prepend-icon="person"
                            name="username"
                            label="Username"
                            type="text"
                            v-model="user.username"
                          ></v-text-field>
                          <v-text-field
                            id="password"
                            prepend-icon="lock"
                            name="password"
                            label="Password"
                            type="password"
                            v-model="user.password"
                          ></v-text-field>
                        </v-form>
                      </v-card-text>
                      <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="secondary" text @click="dialogReg = false">
                          Close
                        </v-btn>
                        <v-btn color="success" @click="registerAccount()">Sign Up</v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                  <Nuxt />
                </div>
              </nav>
              <!-- navbar -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- container -->
      </div>
      <!-- navbar area -->

      <div
        id="home"
        class="header-hero bg_cover"
        style="background-image: url(assets/images/header/banner-bg.svg)"
      >
        <div class="container">
          <div class="row justify-content-center">
            <div class="col-lg-8">
              <div class="header-hero-content text-center">
                <h3
                  class="header-sub-title wow fadeInUp"
                  data-wow-duration="1.3s"
                  data-wow-delay="0.2s"
                >
                  The best way to share your amazing stuff
                </h3>
                <h2
                  class="header-title wow fadeInUp"
                  data-wow-duration="1.3s"
                  data-wow-delay="0.5s"
                >
                 The headless CMS for epic content
                </h2>
                <p
                  class="text wow fadeInUp"
                  data-wow-duration="1.3s"
                  data-wow-delay="0.8s"
                >
                  No more restrictions, no more limits, amazing speed
                </p>
                <a
                  @click="dialogReg = true"
                  class="main-btn wow fadeInUp"
                  data-wow-duration="1.3s"
                  data-wow-delay="1.1s"
                >
                  Get Started
                </a>
              </div>
              <!-- header hero content -->
            </div>
          </div>
          <!-- row -->
          <div class="row">
            <div class="col-lg-12">
              <div
                class="header-hero-image text-center wow fadeIn"
                data-wow-duration="1.3s"
                data-wow-delay="1.4s"
              >
                <img src="assets/images/header/header-hero.png" alt="hero" />
              </div>
              <!-- header hero image -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- container -->
        <div id="particles-1" class="particles"></div>
      </div>
      <!-- header hero -->
    </header>
    <!--====== HEADER PART ENDS ======-->

    <!--====== BRAND PART START ======-->
    <div class="brand-area pt-90" style="display: none">
      <div class="container">
        <div class="row">
          <div class="col-lg-12">
            <div
              class="
                brand-logo
                d-flex
                align-items-center
                justify-content-center justify-content-md-between
              "
            >
              <div
                class="single-logo mt-30 wow fadeIn"
                data-wow-duration="1s"
                data-wow-delay="0.2s"
              >
                <img src="assets/images/brands/uideck.svg" alt="brand" />
              </div>
              <!-- single logo -->
              <div
                class="single-logo mt-30 wow fadeIn"
                data-wow-duration="1.5s"
                data-wow-delay="0.2s"
              >
                <img src="assets/images/brands/ayroui.svg" alt="brand" />
              </div>
              <!-- single logo -->
              <div
                class="single-logo mt-30 wow fadeIn"
                data-wow-duration="1.5s"
                data-wow-delay="0.3s"
              >
                <img src="assets/images/brands/graygrids.svg" alt="brand" />
              </div>
              <!-- single logo -->
              <div
                class="single-logo mt-30 wow fadeIn"
                data-wow-duration="1.5s"
                data-wow-delay="0.4s"
              >
                <img src="assets/images/brands/lineicons.svg" alt="brand" />
              </div>
              <!-- single logo -->
              <div
                class="single-logo mt-30 wow fadeIn"
                data-wow-duration="1.5s"
                data-wow-delay="0.5s"
              >
                <img
                  src="assets/images/brands/ecommerce-html.svg"
                  alt="brand"
                />
              </div>
              <!-- single logo -->
            </div>
            <!-- brand logo -->
          </div>
        </div>
        <!-- row -->
      </div>
      <!-- container -->
    </div>
    <!--====== BRAND PART ENDS ======-->

    <!--====== SERVICES PART START ======-->
    <section id="features" class="services-area pt-120">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-10">
            <div class="section-title text-center pb-40">
              <div class="line m-auto"></div>
              <h3 class="title">
                Clean and simple design,
                <span> Comes with everything you need to get started!</span>
              </h3>
            </div>
            <!-- section title -->
          </div>
        </div>
        <!-- row -->
        <div class="row justify-content-center">
          <div class="col-lg-4 col-md-7 col-sm-8">
            <div
              class="single-services text-center mt-30 wow fadeIn"
              data-wow-duration="1s"
              data-wow-delay="0.2s"
            >
              <div class="services-icon">
                <img
                  class="shape"
                  src="assets/images/services/services-shape.svg"
                  alt="shape"
                />
                <img
                  class="shape-1"
                  src="assets/images/services/services-shape-1.svg"
                  alt="shape"
                />
                <i class="lni lni-baloon"> </i>
              </div>
              <div class="services-content mt-30">
                <h4 class="services-title">
                  <a href="javascript:void(0)">Clean</a>
                </h4>
                <p class="text">
                  Lorem ipsum dolor sit amet,consetetur sadipscing elitr,
                  seddiam nonu eirmod tempor invidunt labore.
                </p>
              </div>
            </div>
            <!-- single services -->
          </div>
          <div class="col-lg-4 col-md-7 col-sm-8">
            <div
              class="single-services text-center mt-30 wow fadeIn"
              data-wow-duration="1s"
              data-wow-delay="0.5s"
            >
              <div class="services-icon">
                <img
                  class="shape"
                  src="assets/images/services/services-shape.svg"
                  alt="shape"
                />
                <img
                  class="shape-1"
                  src="assets/images/services/services-shape-2.svg"
                  alt="shape"
                />
                <i class="lni lni-cog"> </i>
              </div>
              <div class="services-content mt-30">
                <h4 class="services-title">
                  <a href="javascript:void(0)">Robust</a>
                </h4>
                <p class="text">
                  Lorem ipsum dolor sit amet,consetetur sadipscing elitr,
                  seddiam nonu eirmod tempor invidunt labore.
                </p>
              </div>
            </div>
            <!-- single services -->
          </div>
          <div class="col-lg-4 col-md-7 col-sm-8">
            <div
              class="single-services text-center mt-30 wow fadeIn"
              data-wow-duration="1s"
              data-wow-delay="0.8s"
            >
              <div class="services-icon">
                <img
                  class="shape"
                  src="assets/images/services/services-shape.svg"
                  alt="shape"
                />
                <img
                  class="shape-1"
                  src="assets/images/services/services-shape-3.svg"
                  alt="shape"
                />
                <i class="lni lni-bolt-alt"> </i>
              </div>
              <div class="services-content mt-30">
                <h4 class="services-title">
                  <a href="javascript:void(0)">Powerful</a>
                </h4>
                <p class="text">
                  Lorem ipsum dolor sit amet,consetetur sadipscing elitr,
                  seddiam nonu eirmod tempor invidunt labore.
                </p>
              </div>
            </div>
            <!-- single services -->
          </div>
        </div>
        <!-- row -->
      </div>
      <!-- container -->
    </section>
    <!--====== SERVICES PART ENDS ======-->

    <section id="about">
      <!--====== ABOUT PART START ======-->
      <div class="about-area pt-70">
        <div class="container">
          <div class="row">
            <div class="col-lg-6">
              <div
                class="about-content mt-50 wow fadeInLeftBig"
                data-wow-duration="1s"
                data-wow-delay="0.5s"
              >
                <div class="section-title">
                  <div class="line"></div>
                  <h3 class="title">
                    Quick & Easy <span>to Use arualCMS</span>
                  </h3>
                </div>
                <!-- section title -->
                <p class="text">
                  Show your stuff to the whole internet not only to your mum or your friends.
                  We love making good content. In this moment arualCMS is used by freelance
                  developers who are not famous but that want to share their works to the world.
                  Unfortunately with other CMS is creativity blocked or expensive.
                </p>
                <a @click="dialogReg = true" class="main-btn">Try it Free</a>
              </div>
              <!-- about content -->
            </div>
            <div class="col-lg-6">
              <div
                class="about-image text-center mt-50 wow fadeInRightBig"
                data-wow-duration="1s"
                data-wow-delay="0.5s"
              >
                <img src="assets/images/about/about1.svg" alt="about" />
              </div>
              <!-- about image -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- container -->
        <div class="about-shape-1">
          <img src="assets/images/about/about-shape-1.svg" alt="shape" />
        </div>
      </div>
      <!--====== ABOUT PART ENDS ======-->

      <!--====== ABOUT PART START ======-->
      <div class="about-area pt-70">
        <div class="about-shape-2">
          <img src="assets/images/about/about-shape-2.svg" alt="shape" />
        </div>
        <div class="container">
          <div class="row align-items-center">
            <div class="col-lg-6 order-lg-last">
              <div
                class="about-content ms-lg-auto mt-50 wow fadeInLeftBig"
                data-wow-duration="1s"
                data-wow-delay="0.5s"
              >
                <div class="section-title">
                  <div class="line"></div>
                  <h3 class="title">
                    Fast communication <span> with Essential Features</span>
                  </h3>
                </div>
                <!-- section title -->
                <p class="text">
                  Time is important, we don't want you to waste it. So our Rest API
                  reponse is only 30ms. Here you can get a massive influence on real
                  users in minutes. Our CMS give ypu a tool with speed and easy connection.
                  And if your cntent is appreciated you won't only get positive feedback
                  but also lovely and sincere fans
                </p>
                <a @click="dialogReg = true" class="main-btn">Try it Free</a>
              </div>
              <!-- about content -->
            </div>
            <div class="col-lg-6 order-lg-first">
              <div
                class="about-image text-center mt-50 wow fadeInRightBig"
                data-wow-duration="1s"
                data-wow-delay="0.5s"
              >
                <img src="assets/images/about/about2.svg" alt="about" />
              </div>
              <!-- about image -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- container -->
      </div>
      <!--====== ABOUT PART ENDS ======-->

      <!--====== ABOUT PART START ======-->
      <div class="about-area pt-70">
        <div class="container">
          <div class="row">
            <div class="col-lg-6">
              <div
                class="about-content mt-50 wow fadeInLeftBig"
                data-wow-duration="1s"
                data-wow-delay="0.5s"
              >
                <div class="section-title">
                  <div class="line"></div>
                  <h3 class="title">
                    <span>Crafted for</span> SaaS, App and Software Page
                  </h3>
                </div>
                <!-- section title -->
                <p class="text">
                  Imagine if you can directly speak with the world's population.
                  Don't you think it would be easier to find nice people to interact with?
                  arualCMS is both local and global and help you to connect without limitations
                  with people from your city, your state and your universe!
                </p>
                <a @click="dialogReg = true" class="main-btn">Try it Free</a>
              </div>
              <!-- about content -->
            </div>
            <div class="col-lg-6">
              <div
                class="about-image text-center mt-50 wow fadeInRightBig"
                data-wow-duration="1s"
                data-wow-delay="0.5s"
              >
                <img src="assets/images/about/about3.svg" alt="about" />
              </div>
              <!-- about image -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- container -->
        <div class="about-shape-1">
          <img src="assets/images/about/about-shape-1.svg" alt="shape" />
        </div>
      </div>
      <!--====== ABOUT PART ENDS ======-->
    </section>

    <!--====== FOOTER PART START ======-->
    <footer id="footer" class="footer-area pt-120">
      <div class="container">
        <div
          class="subscribe-area wow fadeIn"
          data-wow-duration="1s"
          data-wow-delay="0.5s"
        >
          <div class="row">
            <div class="col-lg-6">
              <div class="subscribe-content mt-45">
                <h2 class="subscribe-title">
                  Subscribe Our Newsletter <span>get regular updates</span>
                </h2>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="subscribe-form mt-50">
                <form action="#">
                  <input type="text" placeholder="Enter eamil" />
                  <button class="main-btn">Subscribe</button>
                </form>
              </div>
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- subscribe area -->
        <div class="footer-widget pb-100">
          <div class="row">
            <div class="col-lg-4 col-md-6 col-sm-8">
              <div
                class="footer-about mt-50 wow fadeIn"
                data-wow-duration="1s"
                data-wow-delay="0.2s"
              >
                <a class="logo" href="javascript:void(0)">
                  <img src="/logosm.png" alt="arualCMS" class="logo"> {{  title }}
                </a>
                <p class="text">
                  We are not a company. We hate companies.
                  Just imagine us like the guys from the Silicon Valley series.
                  We are start up. We made it with love.
                </p>
                <ul class="social">
                  <li>
                    <a href="javascript:void(0)">
                      <i class="lni lni-facebook-filled"> </i>
                    </a>
                  </li>
                  <li>
                    <a href="javascript:void(0)">
                      <i class="lni lni-twitter-filled"> </i>
                    </a>
                  </li>
                  <li>
                    <a href="javascript:void(0)">
                      <i class="lni lni-instagram-filled"> </i>
                    </a>
                  </li>
                  <li>
                    <a href="javascript:void(0)">
                      <i class="lni lni-linkedin-original"> </i>
                    </a>
                  </li>
                </ul>
              </div>
              <!-- footer about -->
            </div>
            <div class="col-lg-5 col-md-7 col-sm-12">
              <div class="footer-link d-flex mt-50 justify-content-sm-between">
                <div
                  class="link-wrapper wow fadeIn"
                  data-wow-duration="1s"
                  data-wow-delay="0.4s"
                >
                  <div class="footer-title">
                    <h4 class="title">Quick Link</h4>
                  </div>
                  <ul class="link">
                    <li><a href="javascript:void(0)">Road Map</a></li>
                    <li><a href="javascript:void(0)">Privacy Policy</a></li>
                    <li><a href="javascript:void(0)">Refund Policy</a></li>
                    <li><a href="javascript:void(0)">Terms of Service</a></li>
                    <li><a href="javascript:void(0)">Pricing</a></li>
                  </ul>
                </div>
                <!-- footer wrapper -->
                <div
                  class="link-wrapper wow fadeIn"
                  data-wow-duration="1s"
                  data-wow-delay="0.6s"
                >
                  <div class="footer-title">
                    <h4 class="title">Resources</h4>
                  </div>
                  <ul class="link">
                    <li><a href="javascript:void(0)">Home</a></li>
                    <li><a href="javascript:void(0)">Page</a></li>
                    <li><a href="javascript:void(0)">Portfolio</a></li>
                    <li><a href="javascript:void(0)">Blog</a></li>
                    <li><a href="javascript:void(0)">Contact</a></li>
                  </ul>
                </div>
                <!-- footer wrapper -->
              </div>
              <!-- footer link -->
            </div>
            <div class="col-lg-3 col-md-5 col-sm-12">
              <div
                class="footer-contact mt-50 wow fadeIn"
                data-wow-duration="1s"
                data-wow-delay="0.8s"
              >
                <div class="footer-title">
                  <h4 class="title">Contact Us</h4>
                </div>
                <ul class="contact">
                  <li>ales.jandera@gmail.com</li>
                  <li>www.ajandera.com</li>
                  <li>
                    Zimna 15, 040 01 Kosice <br />
                    Slovakia.
                  </li>
                </ul>
              </div>
              <!-- footer contact -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- footer widget -->
        <div class="footer-copyright">
          <div class="row">
            <div class="col-lg-12">
              <div class="copyright d-sm-flex justify-content-between">
                <div class="copyright-content">
                  <p class="text">
                    &copy; All right reserved <a href="https://ajandera.com" rel="nofollow">ajandera.com</a>
                  </p>
                </div>
                <!-- copyright content -->
              </div>
              <!-- copyright -->
            </div>
          </div>
          <!-- row -->
        </div>
        <!-- footer copyright -->
      </div>
      <!-- container -->
      <div id="particles-2"></div>
    </footer>
    <!--====== FOOTER PART ENDS ======-->
    <Snackbar />
</v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Snackbar from "~/components/Snackbar.vue";
import ILanding from '~/model/ILanding';
import UserRegistration from '~/model/UserRegitration';
import { namespace } from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component({
  components: {
    Snackbar
  },
  layout: 'sign'
})
export default class SignLayout extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  title: string = "arualCMS";
  register: boolean = false;
  dialogReg: boolean = false;
  imageLink: ILanding = {
    main:
      "/posts.png",
    sub_main:
      "/post-detail.png",
    social_cover:
      "/text-detail.png"
  };
  user: UserRegistration = {username: "", password: "", name: ""}
  $axios: any;

  registerAccount() {
    this.$axios.post("/registration", {
      name: this.user.name,
      username: this.user.username,
      password: this.user.password,
    })
      .then((response: { data: { message: any; success: string }; }) => {
        let res = response.data;
        if (res.success) {
          this.updateText(response.data.message);
          this.updateColor('green');
          this.updateShow(true);
        } else {
          this.updateText(res.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }
}
</script>

<style>

@import "../assets/css/animate.css";
@import "../assets/css/glightbox.min.css";
@import "../assets/css/lineIcons.css";
@import "../assets/css/bootstrap.min.css";
@import "../assets/css/style.css";

.v-parallax__image-container {
  opacity: 0.4;
}
.logo {
  max-height: 60px;
}
.mainHeadline {
  font-weight: 900 !important;
  text-shadow: 3px 2px rgb(204, 202, 202);
  color: #000 !important;
}
.navbar-brand,
.logo {
  color: #fff !important;
}
.signup {
  background: linear-gradient(to right, #fe8464 0%, #fe6e9a 50%, #fe8464 100%);
}
</style>
