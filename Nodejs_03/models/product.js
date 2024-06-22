const products = [];
const path = require("path");

module.exports = class Product {
  constructor(t) {
    this.title = t;
  }

  // function that is a method
  save() {
    products.push(this);
    const p = path.join(
      path.dirname(process.main.filename),
      "data",
      "products.json"
    );
    fs.readFile(p, (err, fileContent) => {
      let existingProducts = [];
      if (!err) {
        existingProducts = JSON.parse(fileContent);
      }
      // because it is an arrow functions this refers to the class
      existingProducts.push(this);
      fs.writeFile(p, JSON.stringify(existingProducts), (err) => {
        console.log(err);
      });
    });
  }

  // call directly in the Product class instead of calling in the created object based on this class
  static fetchAll() {
    // returning the array set outside this class
    return products;
  }
};
