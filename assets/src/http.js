import axios from 'axios';

function handle(res, resolve, reject) {
  if (res.data) {
    if (res.data.error) {
      reject(res.data.error);
    } else {
      resolve(res.data);
    }
  } else {
    console.log(res);
  }
}

class Http {
  constructor() {
    this.domain = process.env.NODE_ENV === 'development'
      ? 'http://localhost:8080'
      : 'https://knowme.theliveup.tv';
  }

  get(path) {
    return new Promise((resolve, reject) => {
      axios.get(`${this.domain}${path}`)
        .then(res => handle(res, resolve, reject))
        .catch(error => console.log(error));
    });
  }

  delete(path) {
    return new Promise((resolve, reject) => {
      axios.delete(`${this.domain}${path}`)
        .then(res => handle(res, resolve, reject))
        .catch(error => console.log(error));
    });
  }

  post(path, data) {
    return new Promise((resolve, reject) => {
      axios.post(`${this.domain}${path}`, data)
        .then(res => handle(res, resolve, reject))
        .catch(error => console.log(error));
    });
  }

  put(path, data) {
    return new Promise((resolve, reject) => {
      axios.put(`${this.domain}${path}`, data)
        .then(res => handle(res, resolve, reject))
        .catch(error => console.log(error));
    });
  }
}

export default new Http();
