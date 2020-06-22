// function timeout(ms) {
//   return new Promise((resolve, reject) => {
//     setTimeout(resolve, ms, "done");
//   });
// }

// timeout(100).then((value) => {
//   console.log(value);
// });

// let promise = new Promise(function (resolve, reject) {
//   console.log("Promise");
//   resolve();
//   console.log("Promise end")
// });

// promise.then(function () {
//   console.log("resolved");
// });

// console.log("HI");

// function loadImageAsync(url) {
//   return new Promise(function (resolve, reject) {
//     const image = new Image();

//     image.onload = function () {
//       resolve(image);
//     };

//     image.onerror = function () {
//       reject(new Error("Could not load image ad " + url));
//     };

//     image.src = url;
//   });
// }

// loadImageAsync("https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png").then((resp) => {
//   console.log(resp)
// })

// const p1 = new Promise(function (resolve, reject) {
//   console.log('p1')
//   setTimeout(() => reject(new Error("fail")), 3000);
// });

// const p2 = new Promise(function (resolve, reject) {
//   console.log('p2')
//   setTimeout(() => {
//     console.log("1s p2");
//     resolve(p1);
//   }, 1000);
// });

// p2.then((result) => console.log(result)).catch((error) => console.log(error));

// getJson

// process.on("unhandledRejection", function (err, p) {
//   throw err;
// });

// const someAsyncThing = function () {
//   return new Promise(function (resolve, reject) {
//     resolve(x + 2);
//   });
// };
// someAsyncThing()
//   .then(function () {
//     console.log("everything is great");
//   })
//   .catch(function (error) {
//     console.log("oh no", error);
//   })
//   .then(function () {
//     console.log("carry on");
//   });
// setTimeout(() => console.log("123"), 2000);

// const promise = new Promise(function (resolve, reject) {
//   resolve("ok");
//   setTimeout(() => {
//     throw new Error("test");
//   }, 0);
// });
// promise.then(function (value) {
//   console.log(value);
// });

// const someAsyncThing = async () => {
//   const promises = [fetch("index.html"), fetch("https://does-not-exist")];
//   const results = await Promise.allSettled(promises);

//   const successfulPromises = results.filter((p) => p.status === "fulfilled");
//   const errors = results
//     .filter((p) => p.status === "rejected")
//     .map((p) => p.reason);
//   console.log(successfulPromises, errors);
// };
// someAsyncThing();

const f = () => console.log("now");
// Promise.resolve().then(f);
(async () => f())();
console.log("next");
