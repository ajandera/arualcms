(function(t){function e(e){for(var n,o,i=e[0],c=e[1],l=e[2],d=0,f=[];d<i.length;d++)o=i[d],Object.prototype.hasOwnProperty.call(s,o)&&s[o]&&f.push(s[o][0]),s[o]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(t[n]=c[n]);u&&u(e);while(f.length)f.shift()();return r.push.apply(r,l||[]),a()}function a(){for(var t,e=0;e<r.length;e++){for(var a=r[e],n=!0,i=1;i<a.length;i++){var c=a[i];0!==s[c]&&(n=!1)}n&&(r.splice(e--,1),t=o(o.s=a[0]))}return t}var n={},s={app:0},r=[];function o(e){if(n[e])return n[e].exports;var a=n[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,o),a.l=!0,a.exports}o.m=t,o.c=n,o.d=function(t,e,a){o.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},o.t=function(t,e){if(1&e&&(t=o(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(o.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)o.d(a,n,function(e){return t[e]}.bind(null,n));return a},o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,"a",e),e},o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},o.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],c=i.push.bind(i);i.push=e,i=i.slice();for(var l=0;l<i.length;l++)e(i[l]);var u=c;r.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},3271:function(t,e,a){"use strict";a("3b30")},"3b30":function(t,e,a){},"56d7":function(t,e,a){"use strict";a.r(e);a("e260"),a("e6cf"),a("cca6"),a("a79d");var n=a("2b0e"),s=a("8c4f"),r=a("2b27"),o=a.n(r),i=a("0628"),c=a.n(i),l=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("Navigation"),a("Header"),a("Content"),a("Footer")],1)},u=[],d=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("nav",{staticClass:"navbar navbar-expand-lg navbar-dark bg-dark fixed-top"},[a("div",{staticClass:"container"},[a("a",{staticClass:"navbar-brand",attrs:{href:"#"}},[t._v(t._s(t.setting.find((function(t){return"title"===t.key})).value))]),t._m(0),t._m(1)])])},f=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("button",{staticClass:"navbar-toggler",attrs:{type:"button","data-toggle":"collapse","data-target":"#navbarResponsive","aria-controls":"navbarResponsive","aria-expanded":"false","aria-label":"Toggle navigation"}},[a("span",{staticClass:"navbar-toggler-icon"})])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"collapse navbar-collapse",attrs:{id:"navbarResponsive"}},[a("ul",{staticClass:"navbar-nav ml-auto"},[a("li",{staticClass:"nav-item active"},[a("a",{staticClass:"nav-link",attrs:{href:"/"}},[t._v("Home "),a("span",{staticClass:"sr-only"},[t._v("(current)")])])]),a("li",{staticClass:"nav-item"},[a("a",{staticClass:"nav-link",attrs:{href:"#news"}},[t._v("News")])]),a("li",{staticClass:"nav-item"},[a("a",{staticClass:"nav-link",attrs:{href:"#contact"}},[t._v("Contact")])])])])}],v=a("bc3a"),p=a.n(v),m={name:"Navigation",components:{},data:function(){return{setting:[]}},mounted:function(){this.load()},methods:{load:function(){var t=this;p.a.get(this.$hostname+"setting").then((function(e){!0===e.data.success?t.setting=e.data.settings:(t.message=e.data.error,t.messageClass="danger")}))}}},h=m,g=a("2877"),b=Object(g["a"])(h,d,f,!1,null,"7c49e56e",null),_=b.exports,C=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("header",{staticClass:"bg-primary py-5 mb-5"},[a("div",{staticClass:"container h-100"},[a("div",{staticClass:"row h-100 align-items-center"},[a("div",{staticClass:"col-lg-12"},[a("h1",{staticClass:"display-4 text-white mt-5 mb-2"},[t._v(t._s(t.setting.find((function(t){return"title"===t.key})).value))]),a("p",{staticClass:"lead mb-5 text-white-50"},[t._v(t._s(t.setting.find((function(t){return"subtitle"===t.key})).value))])])])])])},y=[],w={name:"Header",components:{},data:function(){return{setting:[]}},mounted:function(){this.load()},methods:{load:function(){var t=this;this.posts=[],p.a.get(this.$hostname+"setting").then((function(e){!0===e.data.success?t.setting=e.data.settings:(t.message=e.data.error,t.messageClass="danger")}),(function(t){console.log(t)}))}}},x=w,j=Object(g["a"])(x,C,y,!1,null,"52efbdd9",null),O=j.exports,$=function(){var t=this,e=t.$createElement;t._self._c;return t._m(0)},k=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("footer",{staticClass:"py-5 bg-dark"},[a("div",{staticClass:"container"},[a("p",{staticClass:"m-0 text-center text-white"},[t._v("Copyright © araulCMS by "),a("a",{attrs:{href:"https://ajandera.com"}},[t._v("ajandera.com")]),t._v(" 2021")]),a("p",{staticClass:"m-0 text-center text-white"},[a("a",{attrs:{href:"https://github.com/ajandera/arualcms.git"}},[t._v("Download from github.com")])])])])}],E={name:"Footer",components:{}},P=E,M=Object(g["a"])(P,$,k,!1,null,null,null),H=M.exports,S=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[a("div",{staticClass:"row"},[a("div",{staticClass:"col-md-8 mb-5"},[a("h2",[t._v("What We Do")]),a("hr"),a("div",{domProps:{innerHTML:t._s(t.texts.find((function(t){return"what_we_do"===t.key})).value)}}),a("br"),a("a",{staticClass:"btn btn-primary btn-lg mt-1",attrs:{href:"#"}},[t._v("Call to Action »")])]),a("div",{staticClass:"col-md-4 mb-5",attrs:{id:"contact"}},[a("h2",[t._v("Contact Us")]),a("hr"),a("div",{domProps:{innerHTML:t._s(t.texts.find((function(t){return"contact"===t.key})).value)}})])]),t._m(0),a("div",{staticClass:"row"},t._l(t.posts,(function(e){return a("div",{key:e.id,staticClass:"col-md-4 mb-5"},[a("div",{staticClass:"card h-100"},[a("img",{staticClass:"card-img-top",attrs:{src:"https://placehold.it/300x200",alt:""}}),a("div",{staticClass:"card-body"},[a("h4",{staticClass:"card-title"},[t._v(t._s(e.title))]),a("p",{staticClass:"card-text"},[t._v(t._s(e.excerpt))])]),t._m(1,!0)])])})),0)])},T=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"row",attrs:{id:"news"}},[a("div",{staticClass:"col-md-8 mb-5"},[a("h2",[t._v("News from our blog")]),a("hr")])])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"card-footer"},[a("a",{staticClass:"btn btn-primary",attrs:{href:"#"}},[t._v("Read More!")])])}],N={name:"Content",components:{},data:function(){return{texts:[],posts:[]}},mounted:function(){this.load()},methods:{load:function(){var t=this;p.a.get(this.$hostname+"text").then((function(e){!0===e.data.success?t.texts=e.data.texts:(t.message=e.data.error,t.messageClass="danger")})),p.a.get(this.$hostname+"post").then((function(e){!0===e.data.success?t.posts=e.data.posts:(t.message=e.data.error,t.messageClass="danger")}))}}},R=N,F=Object(g["a"])(R,S,T,!1,null,"22dcf15e",null),A=F.exports,D={name:"App",data:function(){return{message:null,messageClass:null,username:null,password:null,error:null}},components:{Navigation:_,Header:O,Footer:H,Content:A},mounted:function(){},methods:{}},J=D,L=(a("3271"),Object(g["a"])(J,l,u,!1,null,"6112eed0",null)),W=L.exports,U=a("ecee"),q=a("c074"),z=a("ad3d");U["c"].add(q["a"]),n["a"].use(s["a"]),n["a"].use(o.a),n["a"].use(c.a),n["a"].config.productionTip=!0,n["a"].component("font-awesome-icon",z["a"]),n["a"].prototype.$hostname="/api/";var B=[{path:"/",component:W}],G=new s["a"]({mode:"history",routes:B});new n["a"]({router:G,render:function(t){return t(W)}}).$mount("#app")}});
//# sourceMappingURL=app.68855e25.js.map
