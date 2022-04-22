export default function ({ $axios, redirect }) {
  $axios.onError(error => {
    if(error.response.status === 500) {
      redirect('/error')
    }
  })

  $axios.onRequest((config) => {
    const jwt = JSON.parse(window.localStorage.getItem("jwt"));
    config.headers.common['Content-type'] = 'application/json;charset=utf-8';

    // check if the user is authenticated
    if (jwt) {
      // set the Authorization header using the access token
      config.headers.Authorization = 'Bearer ' + jwt.access_token;
    } else {
        redirect('/in')
    }

    return config;
  });
}
