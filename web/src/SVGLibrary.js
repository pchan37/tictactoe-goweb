const SVGHelpers = {
  methods: {
    createSVGElement(type, attributes = {}) {
      const elem = document.createElementNS('http://www.w3.org/2000/svg', type);
      this.setSVGAttributes(elem, attributes);
      return elem;
    },
    createSVGText(text, attributes = {}) {
      const elem = document.createElementNS('http://www.w3.org/2000/svg', 'text');
      const textNode = document.createTextNode(text);
      elem.appendChild(textNode);
      this.setSVGAttributes(elem, attributes);
      return elem;
    },
    setSVGAttributes(element, attributes) {
      Object.keys(attributes).forEach((key) => {
        element.setAttributeNS(null, key, attributes[key]);
      });
    },
  },
};

export default SVGHelpers;
