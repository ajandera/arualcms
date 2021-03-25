(function (t) {
    function e(e) {
        for (var a, r, i = e[0], l = e[1], c = e[2], u = 0, m = []; u < i.length; u++) r = i[u], Object.prototype.hasOwnProperty.call(o, r) && o[r] && m.push(o[r][0]), o[r] = 0;
        for (a in l) Object.prototype.hasOwnProperty.call(l, a) && (t[a] = l[a]);
        d && d(e);
        while (m.length) m.shift()();
        return n.push.apply(n, c || []), s()
    }

    function s() {
        for (var t, e = 0; e < n.length; e++) {
            for (var s = n[e], a = !0, i = 1; i < s.length; i++) {
                var l = s[i];
                0 !== o[l] && (a = !1)
            }
            a && (n.splice(e--, 1), t = r(r.s = s[0]))
        }
        return t
    }

    var a = {}, o = {app: 0}, n = [];

    function r(e) {
        if (a[e]) return a[e].exports;
        var s = a[e] = {i: e, l: !1, exports: {}};
        return t[e].call(s.exports, s, s.exports, r), s.l = !0, s.exports
    }

    r.m = t, r.c = a, r.d = function (t, e, s) {
        r.o(t, e) || Object.defineProperty(t, e, {enumerable: !0, get: s})
    }, r.r = function (t) {
        "undefined" !== typeof Symbol && Symbol.toStringTag && Object.defineProperty(t, Symbol.toStringTag, {value: "Module"}), Object.defineProperty(t, "__esModule", {value: !0})
    }, r.t = function (t, e) {
        if (1 & e && (t = r(t)), 8 & e) return t;
        if (4 & e && "object" === typeof t && t && t.__esModule) return t;
        var s = Object.create(null);
        if (r.r(s), Object.defineProperty(s, "default", {
            enumerable: !0,
            value: t
        }), 2 & e && "string" != typeof t) for (var a in t) r.d(s, a, function (e) {
            return t[e]
        }.bind(null, a));
        return s
    }, r.n = function (t) {
        var e = t && t.__esModule ? function () {
            return t["default"]
        } : function () {
            return t
        };
        return r.d(e, "a", e), e
    }, r.o = function (t, e) {
        return Object.prototype.hasOwnProperty.call(t, e)
    }, r.p = "/";
    var i = window["webpackJsonp"] = window["webpackJsonp"] || [], l = i.push.bind(i);
    i.push = e, i = i.slice();
    for (var c = 0; c < i.length; c++) e(i[c]);
    var d = l;
    n.push([0, "chunk-vendors"]), s()
})({
    0: function (t, e, s) {
        t.exports = s("56d7")
    }, "0cf3": function (t, e, s) {
    }, "1f9d": function (t, e, s) {
        "use strict";
        s("b4c0")
    }, "3b96": function (t, e, s) {
        "use strict";
        s("8bf1")
    }, 4678: function (t, e, s) {
        var a = {
            "./af": "2bfb",
            "./af.js": "2bfb",
            "./ar": "8e73",
            "./ar-dz": "a356",
            "./ar-dz.js": "a356",
            "./ar-kw": "423e",
            "./ar-kw.js": "423e",
            "./ar-ly": "1cfd",
            "./ar-ly.js": "1cfd",
            "./ar-ma": "0a84",
            "./ar-ma.js": "0a84",
            "./ar-sa": "8230",
            "./ar-sa.js": "8230",
            "./ar-tn": "6d83",
            "./ar-tn.js": "6d83",
            "./ar.js": "8e73",
            "./az": "485c",
            "./az.js": "485c",
            "./be": "1fc1",
            "./be.js": "1fc1",
            "./bg": "84aa",
            "./bg.js": "84aa",
            "./bm": "a7fa",
            "./bm.js": "a7fa",
            "./bn": "9043",
            "./bn-bd": "9686",
            "./bn-bd.js": "9686",
            "./bn.js": "9043",
            "./bo": "d26a",
            "./bo.js": "d26a",
            "./br": "6887",
            "./br.js": "6887",
            "./bs": "2554",
            "./bs.js": "2554",
            "./ca": "d716",
            "./ca.js": "d716",
            "./cs": "3c0d",
            "./cs.js": "3c0d",
            "./cv": "03ec",
            "./cv.js": "03ec",
            "./cy": "9797",
            "./cy.js": "9797",
            "./da": "0f14",
            "./da.js": "0f14",
            "./de": "b469",
            "./de-at": "b3eb",
            "./de-at.js": "b3eb",
            "./de-ch": "bb71",
            "./de-ch.js": "bb71",
            "./de.js": "b469",
            "./dv": "598a",
            "./dv.js": "598a",
            "./el": "8d47",
            "./el.js": "8d47",
            "./en-au": "0e6b",
            "./en-au.js": "0e6b",
            "./en-ca": "3886",
            "./en-ca.js": "3886",
            "./en-gb": "39a6",
            "./en-gb.js": "39a6",
            "./en-ie": "e1d3",
            "./en-ie.js": "e1d3",
            "./en-il": "7333",
            "./en-il.js": "7333",
            "./en-in": "ec2e",
            "./en-in.js": "ec2e",
            "./en-nz": "6f50",
            "./en-nz.js": "6f50",
            "./en-sg": "b7e9",
            "./en-sg.js": "b7e9",
            "./eo": "65db",
            "./eo.js": "65db",
            "./es": "898b",
            "./es-do": "0a3c",
            "./es-do.js": "0a3c",
            "./es-mx": "b5b7",
            "./es-mx.js": "b5b7",
            "./es-us": "55c9",
            "./es-us.js": "55c9",
            "./es.js": "898b",
            "./et": "ec18",
            "./et.js": "ec18",
            "./eu": "0ff2",
            "./eu.js": "0ff2",
            "./fa": "8df4",
            "./fa.js": "8df4",
            "./fi": "81e9",
            "./fi.js": "81e9",
            "./fil": "d69a",
            "./fil.js": "d69a",
            "./fo": "0721",
            "./fo.js": "0721",
            "./fr": "9f26",
            "./fr-ca": "d9f8",
            "./fr-ca.js": "d9f8",
            "./fr-ch": "0e49",
            "./fr-ch.js": "0e49",
            "./fr.js": "9f26",
            "./fy": "7118",
            "./fy.js": "7118",
            "./ga": "5120",
            "./ga.js": "5120",
            "./gd": "f6b4",
            "./gd.js": "f6b4",
            "./gl": "8840",
            "./gl.js": "8840",
            "./gom-deva": "aaf2",
            "./gom-deva.js": "aaf2",
            "./gom-latn": "0caa",
            "./gom-latn.js": "0caa",
            "./gu": "e0c5",
            "./gu.js": "e0c5",
            "./he": "c7aa",
            "./he.js": "c7aa",
            "./hi": "dc4d",
            "./hi.js": "dc4d",
            "./hr": "4ba9",
            "./hr.js": "4ba9",
            "./hu": "5b14",
            "./hu.js": "5b14",
            "./hy-am": "d6b6",
            "./hy-am.js": "d6b6",
            "./id": "5038",
            "./id.js": "5038",
            "./is": "0558",
            "./is.js": "0558",
            "./it": "6e98",
            "./it-ch": "6f12",
            "./it-ch.js": "6f12",
            "./it.js": "6e98",
            "./ja": "079e",
            "./ja.js": "079e",
            "./jv": "b540",
            "./jv.js": "b540",
            "./ka": "201b",
            "./ka.js": "201b",
            "./kk": "6d79",
            "./kk.js": "6d79",
            "./km": "e81d",
            "./km.js": "e81d",
            "./kn": "3e92",
            "./kn.js": "3e92",
            "./ko": "22f8",
            "./ko.js": "22f8",
            "./ku": "2421",
            "./ku.js": "2421",
            "./ky": "9609",
            "./ky.js": "9609",
            "./lb": "440c",
            "./lb.js": "440c",
            "./lo": "b29d",
            "./lo.js": "b29d",
            "./lt": "26f9",
            "./lt.js": "26f9",
            "./lv": "b97c",
            "./lv.js": "b97c",
            "./me": "293c",
            "./me.js": "293c",
            "./mi": "688b",
            "./mi.js": "688b",
            "./mk": "6909",
            "./mk.js": "6909",
            "./ml": "02fb",
            "./ml.js": "02fb",
            "./mn": "958b",
            "./mn.js": "958b",
            "./mr": "39bd",
            "./mr.js": "39bd",
            "./ms": "ebe4",
            "./ms-my": "6403",
            "./ms-my.js": "6403",
            "./ms.js": "ebe4",
            "./mt": "1b45",
            "./mt.js": "1b45",
            "./my": "8689",
            "./my.js": "8689",
            "./nb": "6ce3",
            "./nb.js": "6ce3",
            "./ne": "3a39",
            "./ne.js": "3a39",
            "./nl": "facd",
            "./nl-be": "db29",
            "./nl-be.js": "db29",
            "./nl.js": "facd",
            "./nn": "b84c",
            "./nn.js": "b84c",
            "./oc-lnc": "167b",
            "./oc-lnc.js": "167b",
            "./pa-in": "f3ff",
            "./pa-in.js": "f3ff",
            "./pl": "8d57",
            "./pl.js": "8d57",
            "./pt": "f260",
            "./pt-br": "d2d4",
            "./pt-br.js": "d2d4",
            "./pt.js": "f260",
            "./ro": "972c",
            "./ro.js": "972c",
            "./ru": "957c",
            "./ru.js": "957c",
            "./sd": "6784",
            "./sd.js": "6784",
            "./se": "ffff",
            "./se.js": "ffff",
            "./si": "eda5",
            "./si.js": "eda5",
            "./sk": "7be6",
            "./sk.js": "7be6",
            "./sl": "8155",
            "./sl.js": "8155",
            "./sq": "c8f3",
            "./sq.js": "c8f3",
            "./sr": "cf1e",
            "./sr-cyrl": "13e9",
            "./sr-cyrl.js": "13e9",
            "./sr.js": "cf1e",
            "./ss": "52bd",
            "./ss.js": "52bd",
            "./sv": "5fbd",
            "./sv.js": "5fbd",
            "./sw": "74dc",
            "./sw.js": "74dc",
            "./ta": "3de5",
            "./ta.js": "3de5",
            "./te": "5cbb",
            "./te.js": "5cbb",
            "./tet": "576c",
            "./tet.js": "576c",
            "./tg": "3b1b",
            "./tg.js": "3b1b",
            "./th": "10e8",
            "./th.js": "10e8",
            "./tk": "5aff",
            "./tk.js": "5aff",
            "./tl-ph": "0f38",
            "./tl-ph.js": "0f38",
            "./tlh": "cf75",
            "./tlh.js": "cf75",
            "./tr": "0e81",
            "./tr.js": "0e81",
            "./tzl": "cf51",
            "./tzl.js": "cf51",
            "./tzm": "c109",
            "./tzm-latn": "b53d",
            "./tzm-latn.js": "b53d",
            "./tzm.js": "c109",
            "./ug-cn": "6117",
            "./ug-cn.js": "6117",
            "./uk": "ada2",
            "./uk.js": "ada2",
            "./ur": "5294",
            "./ur.js": "5294",
            "./uz": "2e8c",
            "./uz-latn": "010e",
            "./uz-latn.js": "010e",
            "./uz.js": "2e8c",
            "./vi": "2921",
            "./vi.js": "2921",
            "./x-pseudo": "fd7e",
            "./x-pseudo.js": "fd7e",
            "./yo": "7f33",
            "./yo.js": "7f33",
            "./zh-cn": "5c3a",
            "./zh-cn.js": "5c3a",
            "./zh-hk": "49ab",
            "./zh-hk.js": "49ab",
            "./zh-mo": "3a6c",
            "./zh-mo.js": "3a6c",
            "./zh-tw": "90ea",
            "./zh-tw.js": "90ea"
        };

        function o(t) {
            var e = n(t);
            return s(e)
        }

        function n(t) {
            if (!s.o(a, t)) {
                var e = new Error("Cannot find module '" + t + "'");
                throw e.code = "MODULE_NOT_FOUND", e
            }
            return a[t]
        }

        o.keys = function () {
            return Object.keys(a)
        }, o.resolve = n, t.exports = o, o.id = "4678"
    }, "4ca2": function (t, e, s) {
    }, "56d7": function (t, e, s) {
        "use strict";
        s.r(e);
        s("4de4"), s("e260"), s("e6cf"), s("cca6"), s("a79d");
        var a = s("2b0e"), o = s("2ead"), n = s.n(o), r = s("8c4f"), i = s("2b27"), l = s.n(i), c = s("0628"),
            d = s.n(c), u = s("1881"), m = s.n(u), g = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {attrs: {id: "app"}}, [s("section", {staticClass: "main"}, [s("div", {staticClass: "container-fluid"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-2 padding-0 d-none d-sm-block"}, [s("div", {staticClass: "page-wrapper"}, [s("nav", {
                    staticClass: "sidebar-wrapper",
                    attrs: {id: "sidebar"}
                }, [s("div", {staticClass: "sidebar-content"}, [s("div", {staticClass: "sidebar-header"}, [t._m(0), s("div", {staticClass: "user-info"}, [s("span", {staticClass: "user-name"}, [t._v(t._s(t.loggedUser))]), s("span", {staticClass: "user-role"}, [t._v("Administrator")]), s("a", {
                    attrs: {href: "#"},
                    on: {
                        click: function (e) {
                            return t.logout()
                        }
                    }
                }, [t._v("Sign out")])])]), s("div", {staticClass: "sidebar-menu"}, [s("ul", [t._m(1), t._l(t.general, (function (e) {
                    return s("li", {
                        key: e.component,
                        staticClass: "sidebar-dropdown"
                    }, [s("router-link", {attrs: {to: e.component}}, [t._v(t._s(e.label))])], 1)
                })), t._m(2), t._l(t.extra, (function (e) {
                    return s("li", {key: e.component}, [s("router-link", {attrs: {to: e.component}}, [t._v(t._s(e.label))])], 1)
                }))], 2)])]), t._m(3)])])]), s("div", {staticClass: "col-xs-10 col-sm-10"}, [s("nav", {
                    staticClass: "navbar navbar-expand-lg navbar-light bg-light",
                    attrs: {id: "header"}
                }, [s("a", {
                    staticClass: "navbar-brand",
                    attrs: {href: "/"}
                }, [t._v("arualCMS")]), s("div", {staticClass: "topnav d-block d-sm-none"}, [!0 === t.menu ? s("div", {attrs: {id: "myLinks"}}, [s("ul", [t._l(t.general, (function (e) {
                    return s("li", {
                        key: e.component,
                        staticClass: "sidebar-dropdown"
                    }, [s("router-link", {attrs: {to: e.component}}, [t._v(t._s(e.label))])], 1)
                })), t._l(t.extra, (function (e) {
                    return s("li", {key: e.component}, [s("router-link", {attrs: {to: e.component}}, [t._v(t._s(e.label))])], 1)
                }))], 2)]) : t._e(), s("a", {
                    staticClass: "icon", attrs: {href: "#"}, on: {
                        click: function (e) {
                            return t.hamburger()
                        }
                    }
                }, [s("font-awesome-icon", {attrs: {icon: ["fas", "bars"]}})], 1)])]), s("router-view")], 1)])])])])
            }, f = [function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "user-pic"}, [s("img", {
                    staticClass: "img-responsive img-rounded",
                    attrs: {
                        src: "https://raw.githubusercontent.com/azouaoui-med/pro-sidebar-template/gh-pages/src/img/user.jpg",
                        alt: "User picture"
                    }
                })])
            }, function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("li", {staticClass: "header-menu"}, [s("span", [t._v("General")])])
            }, function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("li", {staticClass: "header-menu"}, [s("span", [t._v("Extra")])])
            }, function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "sidebar-footer text-center"}, [s("p", [t._v("2021 © "), s("a", {
                    attrs: {
                        href: "https://ajandera.com",
                        target: "_blank"
                    }
                }, [t._v("ajandera.com")])])])
            }], v = {
                name: "App", data: function () {
                    return {
                        general: [{component: "/admin/posts", label: "Posts"}, {
                            component: "/admin/texts",
                            label: "Texts"
                        }, {component: "/admin/files", label: "Files"}],
                        extra: [{component: "/admin/settings", label: "Settings"}, {component: "/admin/users", label: "Users"}],
                        loggedUser: window.localStorage.getItem("user"),
                        message: null,
                        messageClass: null,
                        error: null,
                        menu: !1
                    }
                }, components: {}, mounted: function () {
                    null === this.loggedUser ? this.$router.push("/admin/signin") : this.$router.push("/adminposts")
                }, methods: {
                    logout: function () {
                        window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), this.loggedUser = null, this.$router.push("signin")
                    }, hamburger: function () {
                        this.menu = !this.menu
                    }
                }
            }, p = v, h = (s("dcf7"), s("2877")), b = Object(h["a"])(p, g, f, !1, null, "bfeaf4bc", null), w = b.exports,
            C = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {attrs: {id: "setting"}}, [t._m(0), t._l(t.setting, (function (e, a) {
                    return s("div", {
                        key: a,
                        staticClass: "row mt-2"
                    }, [s("div", {staticClass: "col-3"}, [s("p", [t._v(t._s(e.key))])]), s("div", {staticClass: "col-8"}, [s("input", {
                        directives: [{
                            name: "model",
                            rawName: "v-model",
                            value: e.value,
                            expression: "item.value"
                        }],
                        staticClass: "form-control",
                        attrs: {type: "text"},
                        domProps: {value: e.value},
                        on: {
                            input: function (s) {
                                s.target.composing || t.$set(e, "value", s.target.value)
                            }
                        }
                    })])])
                })), s("div", {staticClass: "row mt-4"}, [s("div", {staticClass: "col-6"}, [t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e()]), s("div", {staticClass: "col-5"}, [s("div", {staticClass: "float-right"}, [s("button", {
                    staticClass: "btn btn-success",
                    on: {click: t.save}
                }, [t._v("Save")])])])])], 2)
            }, _ = [function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "row"}, [s("div", {staticClass: "col-11"}, [s("h1", [t._v("Settings")]), s("hr")])])
            }], j = s("bc3a"), y = s.n(j), k = {
                name: "Settings", components: {}, data: function () {
                    return {setting: [], messageClass: null, message: null, loggedUser: window.localStorage.getItem("user")}
                }, mounted: function () {
                    this.load()
                }, methods: {
                    load: function () {
                        var t = this;
                        y.a.get(this.$hostname + "setting").then((function (e) {
                            !0 === e.data.success ? t.setting = e.data.settings : (t.message = e.data.error, t.messageClass = "danger")
                        }))
                    }, save: function () {
                        var t = this;
                        y.a.put(this.$hostname + "setting", this.setting, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (e) {
                            e.data.success ? (t.message = e.data.message, t.messageClass = "alert alert-success") : (t.message = e.data.error, t.messageClass = "danger")
                        }), (function (e) {
                            401 === e.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), t.loggedUser = !1, window.location.reload())
                        }))
                    }
                }
            }, x = k, S = Object(h["a"])(x, C, _, !1, null, null, null), I = S.exports, $ = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "table-wrap"}, [s("h1", [t._v("Posts")]), s("hr"), t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e(), s("table", {staticClass: "table table-stripped mt-3"}, [s("thead", [s("tr", [s("th", [t._v("#")]), s("th", [t._v("Title")]), t._m(0), s("th", [t._v("Published")]), s("th", [s("button", {
                    staticClass: "btn btn-secondary btn-success float-right",
                    attrs: {type: "button"},
                    on: {
                        click: function (e) {
                            return t.create()
                        }
                    }
                }, [s("i", {staticClass: "fa fa-plus"}), t._v("Create")])])])]), s("tbody", t._l(t.posts, (function (e) {
                    return s("tr", {
                        key: e.id, staticClass: "actRow", on: {
                            click: function (s) {
                                return t.edit(e.id)
                            }
                        }
                    }, [s("td", [t._v(t._s(e.id))]), s("td", [t._v(t._s(e.title))]), s("td", {staticClass: "d-none d-sm-table-cell"}, [t._v(t._s(e.excerpt))]), s("td", [t._v(t._s(t._f("formatDate")(e.published)))]), s("td", {staticClass: "text-right"}, [s("button", {
                        staticClass: "btn btn-secondary btn-danger",
                        attrs: {type: "button"},
                        on: {
                            click: function (s) {
                                return s.stopPropagation(), s.preventDefault(), t.remove(e.id)
                            }
                        }
                    }, [s("i", {staticClass: "fa fa-trash"}), t._v("Delete")])])])
                })), 0)]), s("modal", {
                    attrs: {
                        name: "form",
                        height: "auto",
                        scrollable: !0,
                        resizable: !0,
                        adaptive: !0
                    }
                }, [s("div", {staticClass: "modal-dialog"}, [s("div", {staticClass: "modal-content"}, [s("div", {staticClass: "modal-header"}, [s("h3", {staticClass: "modal-title"}, [t._v(t._s(t.modalTitle))]), s("button", {
                    staticClass: "close",
                    attrs: {type: "button", "data-dismiss": "modal"},
                    on: {click: t.hide}
                }, [t._v("×")])]), s("div", {staticClass: "modal-body"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [t.error ? s("div", {staticClass: "danger"}, [t._v(t._s(t.error))]) : t._e()])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [s("img", {
                    staticClass: "img-fluid",
                    attrs: {src: t.post.src}
                }), s("br"), s("input", {
                    ref: "file", attrs: {type: "file", id: "file"}, on: {
                        change: function (e) {
                            return t.handleFileUpload()
                        }
                    }
                }), s("div", {staticClass: "float-right mt-3"}, [s("button", {
                    staticClass: "btn btn-lg btn-success",
                    on: {
                        click: function (e) {
                            return t.upload()
                        }
                    }
                }, [t._v("Add Cover")])])])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [s("div", {attrs: {id: "editor"}}, [s("label", [t._v("Title")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.post.title,
                        expression: "post.title"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text"},
                    domProps: {value: t.post.title},
                    on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.post, "title", e.target.value)
                        }
                    }
                }), s("hr"), s("label", [t._v("Published")]), s("br"), s("date-picker", {
                    attrs: {
                        lang: t.lang,
                        type: "datetime",
                        "time-picker-options": t.timePickerOptions
                    }, model: {
                        value: t.post.published, callback: function (e) {
                            t.$set(t.post, "published", e)
                        }, expression: "post.published"
                    }
                }), s("hr"), s("label", [t._v("Excerpt")]), s("textarea", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.post.excerpt,
                        expression: "post.excerpt"
                    }], staticClass: "form-control", domProps: {value: t.post.excerpt}, on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.post, "excerpt", e.target.value)
                        }
                    }
                }), s("hr"), s("quill-editor", {
                    ref: "myQuillEditor",
                    attrs: {options: t.editorOption},
                    on: {
                        blur: function (e) {
                            return t.onEditorBlur(e)
                        }, focus: function (e) {
                            return t.onEditorFocus(e)
                        }, ready: function (e) {
                            return t.onEditorReady(e)
                        }
                    },
                    model: {
                        value: t.post.body, callback: function (e) {
                            t.$set(t.post, "body", e)
                        }, expression: "post.body"
                    }
                }), s("label", {staticClass: "mt-4"}, [t._v("Meta title")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.post.meta.title,
                        expression: "post.meta.title"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text"},
                    domProps: {value: t.post.meta.title},
                    on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.post.meta, "title", e.target.value)
                        }
                    }
                }), s("hr"), s("label", [t._v("Meta description")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.post.meta.description,
                        expression: "post.meta.description"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text"},
                    domProps: {value: t.post.meta.description},
                    on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.post.meta, "description", e.target.value)
                        }
                    }
                }), s("hr"), s("label", [t._v("Meta keywords")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.post.meta.keywords,
                        expression: "post.meta.keywords"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text"},
                    domProps: {value: t.post.meta.keywords},
                    on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.post.meta, "keywords", e.target.value)
                        }
                    }
                }), s("hr"), s("div", {staticClass: "float-right mt-3"}, [s("button", {
                    staticClass: "btn btn-lg btn-success",
                    on: {
                        click: function (e) {
                            return t.save(!0)
                        }
                    }
                }, [t._v("Save")])])], 1)])])])])])])], 1)
            }, E = [function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("th", {staticClass: "d-none d-sm-table-cell"}, [s("br"), t._v("Excerpt")])
            }], P = (s("b0c0"), s("9911"), s("a753"), s("8096"), s("14e1"), s("953d")), z = s("ee3d"), T = s.n(z), U = {
                name: "Posts", components: {quillEditor: P["quillEditor"], DatePicker: T.a}, data: function () {
                    return {
                        messageClass: null,
                        message: null,
                        loggedUser: window.localStorage.getItem("user"),
                        posts: [],
                        post: {title: "", body: "", published: "", meta: {title: "", keywords: "", description: ""}},
                        modalTitle: "",
                        error: "",
                        editorOption: {},
                        lang: {
                            days: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
                            months: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
                            pickers: ["next 7 days", "next 30 days", "previous 7 days", "previous 30 days"],
                            placeholder: {date: "Select Date", dateRange: "Select Date Range"}
                        },
                        timePickerOptions: {start: "00:00", step: "00:30", end: "23:30"}
                    }
                }, mounted: function () {
                    this.load()
                }, methods: {
                    edit: function (t) {
                        this.post = this.posts.filter((function (e) {
                            return e.id === t
                        }))[0], this.modalTitle = this.post.title, this.show()
                    }, remove: function (t) {
                        var e = this;
                        y.a.delete(this.$hostname + "post/" + t, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (t) {
                            t.data.success ? (e.message = t.data.message, e.messageClass = "alert alert-success", e.load()) : (e.message = t.data.message, e.messageClass = "alert alert-danger"), e.hide(), setTimeout((function () {
                                e.message = null, e.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                            401 === t.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), e.loggedUser = !1, window.location.reload())
                        }))
                    }, create: function () {
                        this.post = {
                            title: "",
                            body: "",
                            published: "",
                            meta: {title: "", keywords: "", description: ""}
                        }, this.show()
                    }, show: function () {
                        this.$modal.show("form")
                    }, hide: function () {
                        this.$modal.hide("form")
                    }, save: function (t) {
                        var e = this;
                        void 0 !== this.post.id ? y.a.put(this.$hostname + "post/" + this.post.id, this.post, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (s) {
                            s.data.success ? (e.message = s.data.message, e.messageClass = "alert alert-success") : (e.message = s.data.message, e.messageClass = "alert alert-danger"), e.load(), !0 === t && e.hide(), setTimeout((function () {
                                e.message = null, e.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                            401 === t.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), e.loggedUser = !1, window.location.reload())
                        })) : y.a.post(this.$hostname + "post", this.post, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (t) {
                            t.data.success ? (e.message = t.data.message, e.messageClass = "alert alert-success") : (e.message = t.data.message, e.messageClass = "alert alert-danger"), e.load(), e.hide(), setTimeout((function () {
                                e.message = null, e.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                            401 === t.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), e.loggedUser = !1, window.location.reload())
                        }))
                    }, onEditorBlur: function (t) {
                    }, onEditorFocus: function (t) {
                    }, onEditorReady: function (t) {
                    }, onEditorChange: function (t) {
                        t.quill;
                        var e = t.html;
                        t.text;
                        this.content = e
                    }, load: function () {
                        var t = this;
                        this.posts = [], y.a.get(this.$hostname + "post").then((function (e) {
                            if (!0 === e.data.success) for (var s in e.data.posts) e.data.posts[s].published = new Date(e.data.posts[s].published), t.posts.push(e.data.posts[s]); else t.message = e.data.error, t.messageClass = "danger"
                        }), (function (t) {
                            console.log(t)
                        }))
                    }, handleFileUpload: function () {
                        this.file = this.$refs.file.files[0]
                    }, upload: function () {
                        var t = this, e = new FormData;
                        e.append("file", this.file), y.a.post(this.$hostname + "files/upload", e, {
                            headers: {
                                "Content-Type": "multipart/form-data",
                                Authorization: "Bearer " + window.localStorage.getItem("jwt")
                            }
                        }).then((function (e) {
                            e.data.success ? (t.post.file = e.data.file.name, t.post.src = e.data.file.link, t.save(!1)) : (t.message = e.data.message, t.messageClass = "alert alert-danger"), setTimeout((function () {
                                t.message = null, t.messageClass = null
                            }), 2e3)
                        }), (function (e) {
                            401 === e.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), t.loggedUser = !1, window.location.reload())
                        }))
                    }
                }, computed: {
                    editor: function () {
                        return this.$refs.myQuillEditor.quill
                    }
                }
            }, O = U, M = (s("9cb1"), Object(h["a"])(O, $, E, !1, null, "146886af", null)), F = M.exports, N = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {attrs: {id: "texts"}}, [t._m(0), t._l(t.texts, (function (e, a) {
                    return s("div", {
                        key: a,
                        staticClass: "row mt-2"
                    }, [s("div", {staticClass: "col-sm-1 col-xs-12"}, [a === Object.keys(t.texts).length - 1 ? s("button", {
                        staticClass: "btn btn-warning",
                        on: {click: t.add}
                    }, [s("font-awesome-icon", {attrs: {icon: "plus"}})], 1) : t._e()]), s("div", {staticClass: "col-sm-3 col-xs-6"}, [s("input", {
                        directives: [{
                            name: "model",
                            rawName: "v-model",
                            value: e.key,
                            expression: "item.key"
                        }],
                        staticClass: "form-control",
                        attrs: {type: "text"},
                        domProps: {value: e.key},
                        on: {
                            input: function (s) {
                                s.target.composing || t.$set(e, "key", s.target.value)
                            }
                        }
                    })]), s("div", {staticClass: "col-sm-6 col-xs-6"}, [s("textarea", {
                        directives: [{
                            name: "model",
                            rawName: "v-model",
                            value: e.value,
                            expression: "item.value"
                        }], staticClass: "form-control", domProps: {value: e.value}, on: {
                            input: function (s) {
                                s.target.composing || t.$set(e, "value", s.target.value)
                            }
                        }
                    })]), s("div", {staticClass: "col-sm-1 col-xs-12"}, [s("div", {staticClass: "float-right"}, [s("button", {
                        staticClass: "btn btn-success",
                        on: {
                            click: function (s) {
                                return t.openEditor(e)
                            }
                        }
                    }, [s("font-awesome-icon", {attrs: {icon: "edit"}})], 1), s("button", {
                        staticClass: "btn btn-danger",
                        on: {
                            click: function (s) {
                                return t.remove(e)
                            }
                        }
                    }, [s("font-awesome-icon", {attrs: {icon: "times"}})], 1)])])])
                })), s("div", {staticClass: "row mt-4"}, [s("div", {staticClass: "col-6"}, [t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e()]), s("div", {staticClass: "col-5"}, [s("div", {staticClass: "float-right"}, [s("button", {
                    staticClass: "btn btn-success",
                    on: {click: t.save}
                }, [t._v("Save")])])])]), s("modal", {
                    attrs: {
                        name: "form",
                        width: 800,
                        height: "auto",
                        scrollable: !0
                    }
                }, [s("div", {staticClass: "modal-dialog"}, [s("div", {staticClass: "modal-content"}, [s("div", {staticClass: "modal-header"}, [s("h3", {staticClass: "modal-title"}, [t._v(t._s(t.modalTitle))]), s("button", {
                    staticClass: "close",
                    attrs: {type: "button", "data-dismiss": "modal"},
                    on: {click: t.hide}
                }, [t._v("×")])]), s("div", {staticClass: "modal-body"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [t.error ? s("div", {staticClass: "danger"}, [t._v(t._s(t.error))]) : t._e()])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [s("div", {attrs: {id: "editor"}}, [s("quill-editor", {
                    ref: "myQuillEditor",
                    attrs: {options: t.editorOption},
                    on: {
                        blur: function (e) {
                            return t.onEditorBlur(e)
                        }, focus: function (e) {
                            return t.onEditorFocus(e)
                        }, ready: function (e) {
                            return t.onEditorReady(e)
                        }
                    },
                    model: {
                        value: t.text.value, callback: function (e) {
                            t.$set(t.text, "value", e)
                        }, expression: "text.value"
                    }
                }), s("hr"), s("div", {staticClass: "float-right mt-3"}, [s("button", {
                    staticClass: "btn btn-lg btn-default",
                    on: {click: t.hide}
                }, [t._v("Close")])])], 1)])])])])])])], 2)
            }, q = [function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "row"}, [s("div", {staticClass: "col-11"}, [s("h1", [t._v("Texts")]), s("hr")])])
            }], D = (s("c975"), s("a434"), {
                name: "Texts", components: {quillEditor: P["quillEditor"]}, data: function () {
                    return {
                        texts: [],
                        text: {},
                        messageClass: null,
                        message: null,
                        loggedUser: window.localStorage.getItem("user"),
                        editorOption: {},
                        modalTitle: "",
                        error: ""
                    }
                }, mounted: function () {
                    this.load()
                }, methods: {
                    save: function () {
                        var t = this;
                        y.a.put(this.$hostname + "text", this.texts, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (e) {
                            e.data.success ? (t.message = e.data.message, t.messageClass = "alert alert-success") : (t.message = e.data.error, t.messageClass = "danger")
                        }), (function (e) {
                            401 === e.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), t.loggedUser = !1, window.location.reload())
                        }))
                    }, load: function () {
                        var t = this;
                        y.a.get(this.$hostname + "text").then((function (e) {
                            !0 === e.data.success ? t.texts = e.data.texts : (t.message = e.data.error, t.messageClass = "danger")
                        }))
                    }, add: function () {
                        this.texts.push({key: "", value: ""})
                    }, remove: function (t) {
                        var e = this.texts.indexOf(t);
                        this.texts.splice(e, 1)
                    }, openEditor: function (t) {
                        this.text = t, this.modalTitle = t.key, this.show()
                    }, onEditorBlur: function (t) {
                    }, onEditorFocus: function (t) {
                    }, onEditorReady: function (t) {
                    }, onEditorChange: function (t) {
                        t.quill;
                        var e = t.html;
                        t.text;
                        this.content = e
                    }, show: function () {
                        this.$modal.show("form")
                    }, hide: function () {
                        this.$modal.hide("form")
                    }
                }
            }), A = D, B = (s("67d7"), Object(h["a"])(A, N, q, !1, null, "136afb4a", null)), R = B.exports,
            G = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "table-wrap"}, [s("h1", [t._v("Users")]), s("hr"), t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e(), s("table", {staticClass: "table table-stripped mt-3"}, [s("thead", [s("tr", [s("th", [t._v("#")]), s("th", [t._v("Username")]), s("th", [s("button", {
                    staticClass: "btn btn-secondary btn-success float-right",
                    attrs: {type: "button"},
                    on: {
                        click: function (e) {
                            return t.create()
                        }
                    }
                }, [s("i", {staticClass: "fa fa-plus"}), t._v("Create")])])])]), s("tbody", t._l(t.users, (function (e) {
                    return s("tr", {
                        key: e.id, staticClass: "actRow", on: {
                            click: function (s) {
                                return t.edit(e.id)
                            }
                        }
                    }, [s("td", [t._v(t._s(e.id))]), s("td", [t._v(t._s(e.username))]), s("td", {staticClass: "text-right"}, [s("button", {
                        staticClass: "btn btn-secondary btn-danger",
                        attrs: {type: "button"},
                        on: {
                            click: function (s) {
                                return s.stopPropagation(), s.preventDefault(), t.remove(e.id)
                            }
                        }
                    }, [s("i", {staticClass: "fa fa-trash"}), t._v("Delete")])])])
                })), 0)]), s("modal", {
                    attrs: {
                        name: "form",
                        height: "auto",
                        scrollable: !0,
                        resizable: !0,
                        adaptive: !0
                    }
                }, [s("div", {staticClass: "modal-dialog"}, [s("div", {staticClass: "modal-content"}, [s("div", {staticClass: "modal-header"}, [s("h3", {staticClass: "modal-title"}, [t._v(t._s(t.modalTitle))]), s("button", {
                    staticClass: "close",
                    attrs: {type: "button", "data-dismiss": "modal"},
                    on: {click: t.hide}
                }, [t._v("×")])]), s("div", {staticClass: "modal-body"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [t.error ? s("div", {staticClass: "danger"}, [t._v(t._s(t.error))]) : t._e()])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [s("div", {attrs: {id: "editor"}}, [s("label", [t._v("Username")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.user.username,
                        expression: "user.username"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "email", disabled: void 0 !== this.user.id},
                    domProps: {value: t.user.username},
                    on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.user, "username", e.target.value)
                        }
                    }
                }), s("hr"), s("label", [t._v("Password")]), s("br"), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.user.password,
                        expression: "user.password"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password"},
                    domProps: {value: t.user.password},
                    on: {
                        input: function (e) {
                            e.target.composing || t.$set(t.user, "password", e.target.value)
                        }
                    }
                }), s("div", {staticClass: "float-right mt-3"}, [s("button", {
                    staticClass: "btn btn-lg btn-success",
                    on: {click: t.save}
                }, [t._v("Save")])])])])])])])])])], 1)
            }, J = [], Q = {
                name: "Users", components: {}, data: function () {
                    return {
                        messageClass: null,
                        message: null,
                        loggedUser: window.localStorage.getItem("user"),
                        users: [],
                        user: {username: "", password: ""},
                        modalTitle: "",
                        error: ""
                    }
                }, mounted: function () {
                    this.load()
                }, methods: {
                    edit: function (t) {
                        this.user = this.users.filter((function (e) {
                            return e.id === t
                        }))[0], this.modalTitle = this.user.username, this.show()
                    }, remove: function (t) {
                        var e = this;
                        y.a.delete(this.$hostname + "user/" + t, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (t) {
                            t.data.success ? (e.message = t.data.message, e.messageClass = "alert alert-success", e.load()) : (e.message = t.data.message, e.messageClass = "alert alert-danger"), e.hide(), setTimeout((function () {
                                e.message = null, e.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                            401 === t.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), e.loggedUser = !1, window.location.reload())
                        }))
                    }, create: function () {
                        this.user = {username: "", password: ""}, this.show()
                    }, show: function () {
                        this.$modal.show("form")
                    }, hide: function () {
                        this.$modal.hide("form")
                    }, save: function () {
                        var t = this;
                        void 0 !== this.user.id ? y.a.put(this.$hostname + "user/" + this.user.id, this.user, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (e) {
                            e.data.success ? (t.message = e.data.message, t.messageClass = "alert alert-success") : (t.message = e.data.message, t.messageClass = "alert alert-danger"), t.load(), t.hide(), setTimeout((function () {
                                t.message = null, t.messageClass = null
                            }), 2e3)
                        }), (function (e) {
                            401 === e.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), t.loggedUser = !1, window.location.reload())
                        })) : y.a.post(this.$hostname + "user", this.user, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (e) {
                            e.data.success ? (t.message = e.data.message, t.messageClass = "alert alert-success") : (t.message = e.data.message, t.messageClass = "alert alert-danger"), t.load(), t.hide(), setTimeout((function () {
                                t.message = null, t.messageClass = null
                            }), 2e3)
                        }), (function (e) {
                            401 === e.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), t.loggedUser = !1, window.location.reload())
                        }))
                    }, load: function () {
                        var t = this;
                        this.users = [], y.a.get(this.$hostname + "user").then((function (e) {
                            !0 === e.data.success ? t.users = e.data.users : (t.message = e.data.error, t.messageClass = "danger")
                        }))
                    }
                }, computed: {
                    editor: function () {
                        return this.$refs.myQuillEditor.quill
                    }
                }
            }, Y = Q, L = (s("b30e"), Object(h["a"])(Y, G, J, !1, null, "4e074d65", null)), W = L.exports, H = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("div", {staticClass: "table-wrap"}, [s("h1", [t._v("Files")]), s("hr"), t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e(), s("table", {staticClass: "table table-stripped mt-3"}, [s("thead", [s("tr", [s("th", [t._v("#")]), s("th", [t._v("Image")]), s("th", [t._v("Name")]), s("th", [t._v("Gallery")]), s("th", [s("button", {
                    staticClass: "btn btn-secondary btn-success float-right",
                    attrs: {type: "button"},
                    on: {
                        click: function (e) {
                            return t.create()
                        }
                    }
                }, [s("i", {staticClass: "fa fa-plus"}), t._v("Create")])])])]), s("tbody", t._l(t.files, (function (e) {
                    return s("tr", {
                        key: e.id,
                        staticClass: "actRow"
                    }, [s("td", [t._v(t._s(e.id))]), s("td", [s("img", {
                        staticClass: "img-thumbnail",
                        attrs: {src: e.src}
                    })]), s("td", [t._v(t._s(e.name))]), s("td", [s("input", {
                        directives: [{
                            name: "model",
                            rawName: "v-model",
                            value: e.gallery,
                            expression: "file.gallery"
                        }],
                        staticClass: "form-control",
                        attrs: {type: "text", placeholder: "Gallery"},
                        domProps: {value: e.gallery},
                        on: {
                            change: function (s) {
                                return t.saveGallery(e.id, e.gallery)
                            }, input: function (s) {
                                s.target.composing || t.$set(e, "gallery", s.target.value)
                            }
                        }
                    })]), s("td", {staticClass: "text-right"}, [s("button", {
                        staticClass: "btn btn-secondary btn-danger",
                        attrs: {type: "button"},
                        on: {
                            click: function (s) {
                                return s.stopPropagation(), s.preventDefault(), t.remove(e.id)
                            }
                        }
                    }, [s("i", {staticClass: "fa fa-trash"}), t._v("Delete")])])])
                })), 0)]), s("modal", {
                    attrs: {
                        name: "form",
                        height: "auto",
                        scrollable: !0,
                        resizable: !0,
                        adaptive: !0
                    }
                }, [s("div", {staticClass: "modal-dialog"}, [s("div", {staticClass: "modal-content"}, [s("div", {staticClass: "modal-header"}, [s("h3", {staticClass: "modal-title"}, [t._v(t._s(t.modalTitle))]), s("button", {
                    staticClass: "close",
                    attrs: {type: "button", "data-dismiss": "modal"},
                    on: {click: t.hide}
                }, [t._v("×")])]), s("div", {staticClass: "modal-body"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [t.error ? s("div", {staticClass: "danger"}, [t._v(t._s(t.error))]) : t._e()])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [s("div", {attrs: {id: "editor"}}, [s("label", [t._v("File "), s("input", {
                    ref: "file",
                    attrs: {type: "file", id: "file"},
                    on: {
                        change: function (e) {
                            return t.handleFileUpload()
                        }
                    }
                })]), s("div", {staticClass: "float-right mt-3"}, [s("button", {
                    staticClass: "btn btn-lg btn-success",
                    on: {
                        click: function (e) {
                            return t.save()
                        }
                    }
                }, [t._v("Save")])])])])])])])])])], 1)
            }, K = [], V = {
                name: "Files", components: {}, data: function () {
                    return {
                        messageClass: null,
                        message: null,
                        loggedUser: window.localStorage.getItem("user"),
                        files: [],
                        file: "",
                        modalTitle: "",
                        error: ""
                    }
                }, mounted: function () {
                    this.load()
                }, methods: {
                    remove: function (t) {
                        var e = this;
                        y.a.delete(this.$hostname + "files/" + t, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (t) {
                            t.data.success ? (e.message = t.data.message, e.messageClass = "alert alert-success", e.load()) : (e.message = t.data.message, e.messageClass = "alert alert-danger"), setTimeout((function () {
                                e.message = null, e.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                            401 === t.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), e.loggedUser = !1, window.location.reload())
                        }))
                    }, create: function () {
                        this.show()
                    }, show: function () {
                        this.$modal.show("form")
                    }, hide: function () {
                        this.$modal.hide("form")
                    }, save: function () {
                        var t = this, e = new FormData;
                        e.append("file", this.file), y.a.post(this.$hostname + "files/upload", e, {
                            headers: {
                                "Content-Type": "multipart/form-data",
                                Authorization: "Bearer " + window.localStorage.getItem("jwt")
                            }
                        }).then((function (e) {
                            e.data.success ? (t.message = e.data.message, t.messageClass = "alert alert-success", t.load()) : (t.message = e.data.message, t.messageClass = "alert alert-danger"), t.hide(), setTimeout((function () {
                                t.message = null, t.messageClass = null
                            }), 2e3)
                        }), (function (e) {
                            401 === e.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), t.loggedUser = !1, window.location.reload())
                        }))
                    }, handleFileUpload: function () {
                        this.file = this.$refs.file.files[0]
                    }, load: function () {
                        var t = this;
                        y.a.get(this.$hostname + "files").then((function (e) {
                            !0 === e.data.success ? t.files = e.data.files : (t.message = e.data.error, t.messageClass = "danger")
                        }))
                    }, saveGallery: function (t, e) {
                        var s = this;
                        y.a.put(this.$hostname + "files/gallery/" + t, {gallery: e}, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (t) {
                            t.data.success ? (s.message = t.data.message, s.messageClass = "alert alert-success") : (s.message = t.data.message, s.messageClass = "alert alert-danger"), setTimeout((function () {
                                s.message = null, s.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                            401 === t.response.status && (window.localStorage.removeItem("userId"), window.localStorage.removeItem("user"), window.localStorage.removeItem("jwt"), s.loggedUser = !1, window.location.reload())
                        }))
                    }
                }
            }, X = V, Z = (s("a69e"), Object(h["a"])(X, H, K, !1, null, "0d24c515", null)), tt = Z.exports, et = s("ecee"),
            st = s("c074"), at = s("ad3d"), ot = s("c1df"), nt = s.n(ot), rt = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("modal", {
                    attrs: {
                        name: "sign",
                        height: "auto",
                        clickToClose: !1
                    }
                }, [s("div", {staticClass: "modal-dialog"}, [s("div", {staticClass: "modal-content"}, [s("div", {staticClass: "modal-header"}, [s("h3", {staticClass: "modal-title"}, [t._v("arualCMS")])]), s("div", {staticClass: "modal-body"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e()])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [!1 === t.forgot ? s("div", {attrs: {id: "signin"}}, [s("label", {attrs: {for: "username"}}, [t._v("Username")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.username,
                        expression: "username"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "email", id: "username", required: ""},
                    domProps: {value: t.username},
                    on: {
                        input: function (e) {
                            e.target.composing || (t.username = e.target.value)
                        }
                    }
                }), s("label", {attrs: {for: "password"}}, [t._v("Password")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.password,
                        expression: "password"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password", id: "password", required: ""},
                    domProps: {value: t.password},
                    on: {
                        input: function (e) {
                            e.target.composing || (t.password = e.target.value)
                        }
                    }
                }), s("a", {
                    attrs: {href: "#"}, on: {
                        click: function (e) {
                            return t.forgotForm()
                        }
                    }
                }, [t._v("Forgot your password?")]), s("button", {
                    staticClass: "float-right btn btn-success mt-3",
                    attrs: {type: "submit"},
                    on: {
                        click: function (e) {
                            return t.login()
                        }
                    }
                }, [t._v("Sign In")])]) : t._e(), !0 === t.forgot ? s("div", {attrs: {id: "forgot"}}, [s("label", {attrs: {for: "forgotMail"}}, [t._v("Username")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.forgotMail,
                        expression: "forgotMail"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "email", id: "forgotMail", required: ""},
                    domProps: {value: t.forgotMail},
                    on: {
                        input: function (e) {
                            e.target.composing || (t.forgotMail = e.target.value)
                        }
                    }
                }), s("a", {
                    attrs: {href: "#"}, on: {
                        click: function (e) {
                            return t.forgotForm()
                        }
                    }
                }, [t._v(" back to login")]), s("button", {
                    staticClass: "float-right btn btn-success mt-3",
                    attrs: {type: "submit"},
                    on: {
                        click: function (e) {
                            return t.requestNewPw()
                        }
                    }
                }, [t._v("Request new Password")])]) : t._e()])])])])])])
            }, it = [], lt = {
                name: "Sign", components: {}, data: function () {
                    return {messageClass: null, message: null, forgot: !1, forgotMail: null, username: null, password: null}
                }, mounted: function () {
                    this.$modal.show("sign")
                }, methods: {
                    login: function () {
                        var t = this;
                        y.a.post(this.$hostname + "auth", {
                            username: this.username,
                            password: this.password
                        }).then((function (e) {
                            var s = e.data;
                            t.message = s.message, s.success ? (t.message = s.message, t.messageClass = "alert alert-success", t.username = null, t.password = null, t.loggedUser = s.username, window.localStorage.setItem("userId", s.id), window.localStorage.setItem("user", s.username), window.localStorage.setItem("jwt", s.jwt), t.hide(), t.$router.push("/admin/posts")) : (t.message = s.message, t.messageClass = "alert alert-danger")
                        }))
                    }, hide: function () {
                        this.$modal.hide("sign")
                    }, forgotForm: function () {
                        this.forgot = !this.forgot
                    }, requestNewPw: function () {
                        var t = this,
                            e = window.location.protocol + "//" + window.location.hostname + "/recovery/" + this.forgotMail,
                            s = {
                                email: this.forgotMail,
                                subject: "Recovery your password",
                                body: "For recovery your password use this link: <a href='" + e + "'>" + e + "</a>"
                            };
                        y.a.post(this.$hostname + "mail", s, {headers: {Authorization: "Bearer " + window.localStorage.getItem("jwt")}}).then((function (e) {
                            var s = e.data;
                            t.message = s.message, s.success ? (t.message = s.message, t.messageClass = "alert alert-success", t.forgotMail = null, t.forgot()) : (t.message = s.message, t.messageClass = "alert alert-danger")
                        }))
                    }
                }
            }, ct = lt, dt = (s("1f9d"), Object(h["a"])(ct, rt, it, !1, null, "c011d318", null)), ut = dt.exports,
            mt = function () {
                var t = this, e = t.$createElement, s = t._self._c || e;
                return s("modal", {
                    attrs: {
                        name: "recovery",
                        height: "auto",
                        clickToClose: !1
                    }
                }, [s("div", {staticClass: "modal-dialog"}, [s("div", {staticClass: "modal-content"}, [s("div", {staticClass: "modal-header"}, [s("h3", {staticClass: "modal-title"}, [t._v("arualCMS")])]), s("div", {staticClass: "modal-body"}, [s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [t.message ? s("div", {class: t.messageClass}, [t._v(t._s(t.message))]) : t._e()])]), s("div", {staticClass: "row"}, [s("div", {staticClass: "col-12"}, [s("div", {attrs: {id: "signin"}}, [s("label", {attrs: {for: "password"}}, [t._v("Password " + t._s(t.$route.params.username))]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.password,
                        expression: "password"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password", id: "password", required: ""},
                    domProps: {value: t.password},
                    on: {
                        input: function (e) {
                            e.target.composing || (t.password = e.target.value)
                        }
                    }
                }), s("label", {attrs: {for: "passwordCheck"}}, [t._v("Password check "), t.passwordCheck.length > 0 && t.password !== t.passwordCheck ? s("span", {staticClass: "red"}, [t._v("(Password not match!)")]) : t._e()]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.passwordCheck,
                        expression: "passwordCheck"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password", id: "passwordCheck", required: ""},
                    domProps: {value: t.passwordCheck},
                    on: {
                        input: function (e) {
                            e.target.composing || (t.passwordCheck = e.target.value)
                        }
                    }
                }), s("button", {
                    staticClass: "float-right btn btn-success mt-3",
                    attrs: {type: "submit", disabled: t.password !== t.passwordCheck},
                    on: {
                        click: function (e) {
                            return t.recoveryPw()
                        }
                    }
                }, [t._v("Set password")])])])])])])])])
            }, gt = [], ft = {
                name: "Recovery", components: {}, data: function () {
                    return {messageClass: null, message: null, password: null, passwordCheck: ""}
                }, mounted: function () {
                    this.$modal.show("recovery")
                }, methods: {
                    recoveryPw: function () {
                        var t = this, e = this.$route.params.username;
                        y.a.post(this.$hostname + "recovery", {username: e, password: this.password}).then((function (e) {
                            e.data.success ? (t.message = e.data.message, t.messageClass = "alert alert-success") : (t.message = e.data.message, t.messageClass = "alert alert-danger"), setTimeout((function () {
                                t.message = null, t.messageClass = null
                            }), 2e3)
                        }), (function (t) {
                        }))
                    }
                }
            }, vt = ft, pt = (s("3b96"), Object(h["a"])(vt, mt, gt, !1, null, "39aec8cf", null)), ht = pt.exports;
        et["c"].add(st["a"]), a["a"].use(n.a), a["a"].use(r["a"]), a["a"].use(l.a), a["a"].use(d.a), a["a"].use(m.a), a["a"].config.productionTip = !0, a["a"].component("font-awesome-icon", at["a"]), a["a"].prototype.$hostname = "/api/", a["a"].filter("formatDate", (function (t) {
            if (t) return nt()(String(t)).format("DD/MM/YYYY hh:mm")
        }));
        var bt = [{path: "/", component: w}, {path: "/admin/posts", component: F}, {
                path: "/admin/files",
                component: tt
            }, {path: "/admin/settings", component: I}, {path: "/admin/texts", component: R}, {
                path: "/admin/users",
                component: W
            }, {path: "/admin/signin", component: ut}, {path: "/admin/recovery/:username", component: ht}],
            wt = new r["a"]({mode: "history", routes: bt});
        new a["a"]({
            router: wt, render: function (t) {
                return t(w)
            }
        }).$mount("#app")
    }, "59c9": function (t, e, s) {
    }, "67d7": function (t, e, s) {
        "use strict";
        s("e3cf")
    }, "8bf1": function (t, e, s) {
    }, "9cb1": function (t, e, s) {
        "use strict";
        s("59c9")
    }, a69e: function (t, e, s) {
        "use strict";
        s("0cf3")
    }, b30e: function (t, e, s) {
        "use strict";
        s("bd28")
    }, b4c0: function (t, e, s) {
    }, bd28: function (t, e, s) {
    }, dcf7: function (t, e, s) {
        "use strict";
        s("4ca2")
    }, e3cf: function (t, e, s) {
    }
});
//# sourceMappingURL=app.c21797ac.js.map
