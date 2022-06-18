export default function ({ $axios, redirect }) {
  $axios.onError(error => {
    console.log(error);
    if(error.response.status === 500) {
      redirect('/error')
    }
  })

  $axios.onRequest((config) => {
    const jwt = window.localStorage.getItem("jwt");
    config.headers.common['Content-Type'] = "application/json;charset=utf-8";
    if (config.url !== '/auth') {
      // check if the user is authenticated
      if (jwt) {
        // set the Authorization header using the access token
        config.headers.common['Authorization'] = 'Bearer ' + jwt;
      } else {
        redirect('/in')
      }
    }
    return config;
  });
}
