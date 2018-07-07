import axios from 'axios';

class Http {
  constructor() {
    this.domain = process.env.NODE_ENV === 'development'
      ? 'http://localhost:8080'
      : 'http://knowme.theliveup.tv';
  }

  get(path) {
    return new Promise((resolve) => {
      axios.get(`${this.domain}${path}`)
        .then((res) => {
          if (res.data) {
            resolve(res.data);
          } else {
            console.log(res);
          }
        })
        .catch(error => console.log(error));
    });
  }

  delete(path) {
    return new Promise((resolve) => {
      axios.delete(`${this.domain}${path}`)
        .then((res) => {
          if (res.data) {
            resolve(res.data);
          } else {
            console.log(res);
          }
        })
        .catch(error => console.log(error));
    });
  }

  post(path, data) {
    return new Promise((resolve) => {
      axios.post(`${this.domain}${path}`, data)
        .then((res) => {
          if (res.data) {
            resolve(res.data);
          } else {
            console.log(res);
          }
        })
        .catch(error => console.log(error));
    });
  }

  put(path, data) {
    return new Promise((resolve) => {
      axios.put(`${this.domain}${path}`, data)
        .then((res) => {
          if (res.data) {
            resolve(res.data);
          } else {
            console.log(res);
          }
        })
        .catch(error => console.log(error));
    });
  }
}

export default new Http();
