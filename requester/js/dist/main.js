var t = (() => {
  var __defProp = Object.defineProperty;
  var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
  var __markAsModule = (target) => __defProp(target, "__esModule", { value: true });
  var __esm = (fn, res) => function __init() {
    return fn && (res = (0, fn[Object.keys(fn)[0]])(fn = 0)), res;
  };
  var __export = (target, all) => {
    __markAsModule(target);
    for (var name in all)
      __defProp(target, name, { get: all[name], enumerable: true });
  };
  var __publicField = (obj, key, value) => {
    __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
    return value;
  };
  var __accessCheck = (obj, member, msg) => {
    if (!member.has(obj))
      throw TypeError("Cannot " + msg);
  };
  var __privateGet = (obj, member, getter) => {
    __accessCheck(obj, member, "read from private field");
    return getter ? getter.call(obj) : member.get(obj);
  };
  var __privateAdd = (obj, member, value) => {
    if (member.has(obj))
      throw TypeError("Cannot add the same private member more than once");
    member instanceof WeakSet ? member.add(obj) : member.set(obj, value);
  };
  var __privateSet = (obj, member, value, setter) => {
    __accessCheck(obj, member, "write to private field");
    setter ? setter.call(obj, value) : member.set(obj, value);
    return value;
  };

  // node_modules/jsonpath-plus/dist/index-browser-esm.js
  var index_browser_esm_exports = {};
  __export(index_browser_esm_exports, {
    JSONPath: () => JSONPath
  });
  function _typeof(obj) {
    "@babel/helpers - typeof";
    if (typeof Symbol === "function" && typeof Symbol.iterator === "symbol") {
      _typeof = function(obj2) {
        return typeof obj2;
      };
    } else {
      _typeof = function(obj2) {
        return obj2 && typeof Symbol === "function" && obj2.constructor === Symbol && obj2 !== Symbol.prototype ? "symbol" : typeof obj2;
      };
    }
    return _typeof(obj);
  }
  function _classCallCheck(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
      throw new TypeError("Cannot call a class as a function");
    }
  }
  function _inherits(subClass, superClass) {
    if (typeof superClass !== "function" && superClass !== null) {
      throw new TypeError("Super expression must either be null or a function");
    }
    subClass.prototype = Object.create(superClass && superClass.prototype, {
      constructor: {
        value: subClass,
        writable: true,
        configurable: true
      }
    });
    if (superClass)
      _setPrototypeOf(subClass, superClass);
  }
  function _getPrototypeOf(o) {
    _getPrototypeOf = Object.setPrototypeOf ? Object.getPrototypeOf : function _getPrototypeOf2(o2) {
      return o2.__proto__ || Object.getPrototypeOf(o2);
    };
    return _getPrototypeOf(o);
  }
  function _setPrototypeOf(o, p) {
    _setPrototypeOf = Object.setPrototypeOf || function _setPrototypeOf2(o2, p2) {
      o2.__proto__ = p2;
      return o2;
    };
    return _setPrototypeOf(o, p);
  }
  function _isNativeReflectConstruct() {
    if (typeof Reflect === "undefined" || !Reflect.construct)
      return false;
    if (Reflect.construct.sham)
      return false;
    if (typeof Proxy === "function")
      return true;
    try {
      Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], function() {
      }));
      return true;
    } catch (e) {
      return false;
    }
  }
  function _construct(Parent, args, Class) {
    if (_isNativeReflectConstruct()) {
      _construct = Reflect.construct;
    } else {
      _construct = function _construct2(Parent2, args2, Class2) {
        var a = [null];
        a.push.apply(a, args2);
        var Constructor = Function.bind.apply(Parent2, a);
        var instance = new Constructor();
        if (Class2)
          _setPrototypeOf(instance, Class2.prototype);
        return instance;
      };
    }
    return _construct.apply(null, arguments);
  }
  function _isNativeFunction(fn) {
    return Function.toString.call(fn).indexOf("[native code]") !== -1;
  }
  function _wrapNativeSuper(Class) {
    var _cache = typeof Map === "function" ? /* @__PURE__ */ new Map() : void 0;
    _wrapNativeSuper = function _wrapNativeSuper2(Class2) {
      if (Class2 === null || !_isNativeFunction(Class2))
        return Class2;
      if (typeof Class2 !== "function") {
        throw new TypeError("Super expression must either be null or a function");
      }
      if (typeof _cache !== "undefined") {
        if (_cache.has(Class2))
          return _cache.get(Class2);
        _cache.set(Class2, Wrapper);
      }
      function Wrapper() {
        return _construct(Class2, arguments, _getPrototypeOf(this).constructor);
      }
      Wrapper.prototype = Object.create(Class2.prototype, {
        constructor: {
          value: Wrapper,
          enumerable: false,
          writable: true,
          configurable: true
        }
      });
      return _setPrototypeOf(Wrapper, Class2);
    };
    return _wrapNativeSuper(Class);
  }
  function _assertThisInitialized(self) {
    if (self === void 0) {
      throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
    }
    return self;
  }
  function _possibleConstructorReturn(self, call) {
    if (call && (typeof call === "object" || typeof call === "function")) {
      return call;
    }
    return _assertThisInitialized(self);
  }
  function _createSuper(Derived) {
    var hasNativeReflectConstruct = _isNativeReflectConstruct();
    return function _createSuperInternal() {
      var Super = _getPrototypeOf(Derived), result;
      if (hasNativeReflectConstruct) {
        var NewTarget = _getPrototypeOf(this).constructor;
        result = Reflect.construct(Super, arguments, NewTarget);
      } else {
        result = Super.apply(this, arguments);
      }
      return _possibleConstructorReturn(this, result);
    };
  }
  function _toConsumableArray(arr) {
    return _arrayWithoutHoles(arr) || _iterableToArray(arr) || _unsupportedIterableToArray(arr) || _nonIterableSpread();
  }
  function _arrayWithoutHoles(arr) {
    if (Array.isArray(arr))
      return _arrayLikeToArray(arr);
  }
  function _iterableToArray(iter) {
    if (typeof Symbol !== "undefined" && iter[Symbol.iterator] != null || iter["@@iterator"] != null)
      return Array.from(iter);
  }
  function _unsupportedIterableToArray(o, minLen) {
    if (!o)
      return;
    if (typeof o === "string")
      return _arrayLikeToArray(o, minLen);
    var n = Object.prototype.toString.call(o).slice(8, -1);
    if (n === "Object" && o.constructor)
      n = o.constructor.name;
    if (n === "Map" || n === "Set")
      return Array.from(o);
    if (n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))
      return _arrayLikeToArray(o, minLen);
  }
  function _arrayLikeToArray(arr, len) {
    if (len == null || len > arr.length)
      len = arr.length;
    for (var i = 0, arr2 = new Array(len); i < len; i++)
      arr2[i] = arr[i];
    return arr2;
  }
  function _nonIterableSpread() {
    throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.");
  }
  function _createForOfIteratorHelper(o, allowArrayLike) {
    var it = typeof Symbol !== "undefined" && o[Symbol.iterator] || o["@@iterator"];
    if (!it) {
      if (Array.isArray(o) || (it = _unsupportedIterableToArray(o)) || allowArrayLike && o && typeof o.length === "number") {
        if (it)
          o = it;
        var i = 0;
        var F = function() {
        };
        return {
          s: F,
          n: function() {
            if (i >= o.length)
              return {
                done: true
              };
            return {
              done: false,
              value: o[i++]
            };
          },
          e: function(e) {
            throw e;
          },
          f: F
        };
      }
      throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.");
    }
    var normalCompletion = true, didErr = false, err;
    return {
      s: function() {
        it = it.call(o);
      },
      n: function() {
        var step = it.next();
        normalCompletion = step.done;
        return step;
      },
      e: function(e) {
        didErr = true;
        err = e;
      },
      f: function() {
        try {
          if (!normalCompletion && it.return != null)
            it.return();
        } finally {
          if (didErr)
            throw err;
        }
      }
    };
  }
  function push(arr, item) {
    arr = arr.slice();
    arr.push(item);
    return arr;
  }
  function unshift(item, arr) {
    arr = arr.slice();
    arr.unshift(item);
    return arr;
  }
  function JSONPath(opts, expr, obj, callback, otherTypeCallback) {
    if (!(this instanceof JSONPath)) {
      try {
        return new JSONPath(opts, expr, obj, callback, otherTypeCallback);
      } catch (e) {
        if (!e.avoidNew) {
          throw e;
        }
        return e.value;
      }
    }
    if (typeof opts === "string") {
      otherTypeCallback = callback;
      callback = obj;
      obj = expr;
      expr = opts;
      opts = null;
    }
    var optObj = opts && _typeof(opts) === "object";
    opts = opts || {};
    this.json = opts.json || obj;
    this.path = opts.path || expr;
    this.resultType = opts.resultType || "value";
    this.flatten = opts.flatten || false;
    this.wrap = hasOwnProp.call(opts, "wrap") ? opts.wrap : true;
    this.sandbox = opts.sandbox || {};
    this.preventEval = opts.preventEval || false;
    this.parent = opts.parent || null;
    this.parentProperty = opts.parentProperty || null;
    this.callback = opts.callback || callback || null;
    this.otherTypeCallback = opts.otherTypeCallback || otherTypeCallback || function() {
      throw new TypeError("You must supply an otherTypeCallback callback option with the @other() operator.");
    };
    if (opts.autostart !== false) {
      var args = {
        path: optObj ? opts.path : expr
      };
      if (!optObj) {
        args.json = obj;
      } else if ("json" in opts) {
        args.json = opts.json;
      }
      var ret = this.evaluate(args);
      if (!ret || _typeof(ret) !== "object") {
        throw new NewError(ret);
      }
      return ret;
    }
  }
  var hasOwnProp, NewError, moveToAnotherArray;
  var init_index_browser_esm = __esm({
    "node_modules/jsonpath-plus/dist/index-browser-esm.js"() {
      hasOwnProp = Object.prototype.hasOwnProperty;
      NewError = /* @__PURE__ */ function(_Error) {
        _inherits(NewError2, _Error);
        var _super = _createSuper(NewError2);
        function NewError2(value) {
          var _this;
          _classCallCheck(this, NewError2);
          _this = _super.call(this, 'JSONPath should not be called with "new" (it prevents return of (unwrapped) scalar values)');
          _this.avoidNew = true;
          _this.value = value;
          _this.name = "NewError";
          return _this;
        }
        return NewError2;
      }(/* @__PURE__ */ _wrapNativeSuper(Error));
      JSONPath.prototype.evaluate = function(expr, json, callback, otherTypeCallback) {
        var _this2 = this;
        var currParent = this.parent, currParentProperty = this.parentProperty;
        var flatten = this.flatten, wrap = this.wrap;
        this.currResultType = this.resultType;
        this.currPreventEval = this.preventEval;
        this.currSandbox = this.sandbox;
        callback = callback || this.callback;
        this.currOtherTypeCallback = otherTypeCallback || this.otherTypeCallback;
        json = json || this.json;
        expr = expr || this.path;
        if (expr && _typeof(expr) === "object" && !Array.isArray(expr)) {
          if (!expr.path && expr.path !== "") {
            throw new TypeError('You must supply a "path" property when providing an object argument to JSONPath.evaluate().');
          }
          if (!hasOwnProp.call(expr, "json")) {
            throw new TypeError('You must supply a "json" property when providing an object argument to JSONPath.evaluate().');
          }
          var _expr = expr;
          json = _expr.json;
          flatten = hasOwnProp.call(expr, "flatten") ? expr.flatten : flatten;
          this.currResultType = hasOwnProp.call(expr, "resultType") ? expr.resultType : this.currResultType;
          this.currSandbox = hasOwnProp.call(expr, "sandbox") ? expr.sandbox : this.currSandbox;
          wrap = hasOwnProp.call(expr, "wrap") ? expr.wrap : wrap;
          this.currPreventEval = hasOwnProp.call(expr, "preventEval") ? expr.preventEval : this.currPreventEval;
          callback = hasOwnProp.call(expr, "callback") ? expr.callback : callback;
          this.currOtherTypeCallback = hasOwnProp.call(expr, "otherTypeCallback") ? expr.otherTypeCallback : this.currOtherTypeCallback;
          currParent = hasOwnProp.call(expr, "parent") ? expr.parent : currParent;
          currParentProperty = hasOwnProp.call(expr, "parentProperty") ? expr.parentProperty : currParentProperty;
          expr = expr.path;
        }
        currParent = currParent || null;
        currParentProperty = currParentProperty || null;
        if (Array.isArray(expr)) {
          expr = JSONPath.toPathString(expr);
        }
        if (!expr && expr !== "" || !json) {
          return void 0;
        }
        var exprList = JSONPath.toPathArray(expr);
        if (exprList[0] === "$" && exprList.length > 1) {
          exprList.shift();
        }
        this._hasParentSelector = null;
        var result = this._trace(exprList, json, ["$"], currParent, currParentProperty, callback).filter(function(ea) {
          return ea && !ea.isParentSelector;
        });
        if (!result.length) {
          return wrap ? [] : void 0;
        }
        if (!wrap && result.length === 1 && !result[0].hasArrExpr) {
          return this._getPreferredOutput(result[0]);
        }
        return result.reduce(function(rslt, ea) {
          var valOrPath = _this2._getPreferredOutput(ea);
          if (flatten && Array.isArray(valOrPath)) {
            rslt = rslt.concat(valOrPath);
          } else {
            rslt.push(valOrPath);
          }
          return rslt;
        }, []);
      };
      JSONPath.prototype._getPreferredOutput = function(ea) {
        var resultType = this.currResultType;
        switch (resultType) {
          case "all": {
            var path = Array.isArray(ea.path) ? ea.path : JSONPath.toPathArray(ea.path);
            ea.pointer = JSONPath.toPointer(path);
            ea.path = typeof ea.path === "string" ? ea.path : JSONPath.toPathString(ea.path);
            return ea;
          }
          case "value":
          case "parent":
          case "parentProperty":
            return ea[resultType];
          case "path":
            return JSONPath.toPathString(ea[resultType]);
          case "pointer":
            return JSONPath.toPointer(ea.path);
          default:
            throw new TypeError("Unknown result type");
        }
      };
      JSONPath.prototype._handleCallback = function(fullRetObj, callback, type) {
        if (callback) {
          var preferredOutput = this._getPreferredOutput(fullRetObj);
          fullRetObj.path = typeof fullRetObj.path === "string" ? fullRetObj.path : JSONPath.toPathString(fullRetObj.path);
          callback(preferredOutput, type, fullRetObj);
        }
      };
      JSONPath.prototype._trace = function(expr, val, path, parent, parentPropName, callback, hasArrExpr, literalPriority) {
        var _this3 = this;
        var retObj;
        if (!expr.length) {
          retObj = {
            path,
            value: val,
            parent,
            parentProperty: parentPropName,
            hasArrExpr
          };
          this._handleCallback(retObj, callback, "value");
          return retObj;
        }
        var loc = expr[0], x = expr.slice(1);
        var ret = [];
        function addRet(elems) {
          if (Array.isArray(elems)) {
            elems.forEach(function(t2) {
              ret.push(t2);
            });
          } else {
            ret.push(elems);
          }
        }
        if ((typeof loc !== "string" || literalPriority) && val && hasOwnProp.call(val, loc)) {
          addRet(this._trace(x, val[loc], push(path, loc), val, loc, callback, hasArrExpr));
        } else if (loc === "*") {
          this._walk(loc, x, val, path, parent, parentPropName, callback, function(m, l, _x, v, p, par, pr, cb) {
            addRet(_this3._trace(unshift(m, _x), v, p, par, pr, cb, true, true));
          });
        } else if (loc === "..") {
          addRet(this._trace(x, val, path, parent, parentPropName, callback, hasArrExpr));
          this._walk(loc, x, val, path, parent, parentPropName, callback, function(m, l, _x, v, p, par, pr, cb) {
            if (_typeof(v[m]) === "object") {
              addRet(_this3._trace(unshift(l, _x), v[m], push(p, m), v, m, cb, true));
            }
          });
        } else if (loc === "^") {
          this._hasParentSelector = true;
          return {
            path: path.slice(0, -1),
            expr: x,
            isParentSelector: true
          };
        } else if (loc === "~") {
          retObj = {
            path: push(path, loc),
            value: parentPropName,
            parent,
            parentProperty: null
          };
          this._handleCallback(retObj, callback, "property");
          return retObj;
        } else if (loc === "$") {
          addRet(this._trace(x, val, path, null, null, callback, hasArrExpr));
        } else if (/^(\x2D?[0-9]*):(\x2D?[0-9]*):?([0-9]*)$/.test(loc)) {
          addRet(this._slice(loc, x, val, path, parent, parentPropName, callback));
        } else if (loc.indexOf("?(") === 0) {
          if (this.currPreventEval) {
            throw new Error("Eval [?(expr)] prevented in JSONPath expression.");
          }
          this._walk(loc, x, val, path, parent, parentPropName, callback, function(m, l, _x, v, p, par, pr, cb) {
            if (_this3._eval(l.replace(/^\?\(((?:[\0-\t\x0B\f\x0E-\u2027\u202A-\uD7FF\uE000-\uFFFF]|[\uD800-\uDBFF][\uDC00-\uDFFF]|[\uD800-\uDBFF](?![\uDC00-\uDFFF])|(?:[^\uD800-\uDBFF]|^)[\uDC00-\uDFFF])*?)\)$/, "$1"), v[m], m, p, par, pr)) {
              addRet(_this3._trace(unshift(m, _x), v, p, par, pr, cb, true));
            }
          });
        } else if (loc[0] === "(") {
          if (this.currPreventEval) {
            throw new Error("Eval [(expr)] prevented in JSONPath expression.");
          }
          addRet(this._trace(unshift(this._eval(loc, val, path[path.length - 1], path.slice(0, -1), parent, parentPropName), x), val, path, parent, parentPropName, callback, hasArrExpr));
        } else if (loc[0] === "@") {
          var addType = false;
          var valueType = loc.slice(1, -2);
          switch (valueType) {
            case "scalar":
              if (!val || !["object", "function"].includes(_typeof(val))) {
                addType = true;
              }
              break;
            case "boolean":
            case "string":
            case "undefined":
            case "function":
              if (_typeof(val) === valueType) {
                addType = true;
              }
              break;
            case "integer":
              if (Number.isFinite(val) && !(val % 1)) {
                addType = true;
              }
              break;
            case "number":
              if (Number.isFinite(val)) {
                addType = true;
              }
              break;
            case "nonFinite":
              if (typeof val === "number" && !Number.isFinite(val)) {
                addType = true;
              }
              break;
            case "object":
              if (val && _typeof(val) === valueType) {
                addType = true;
              }
              break;
            case "array":
              if (Array.isArray(val)) {
                addType = true;
              }
              break;
            case "other":
              addType = this.currOtherTypeCallback(val, path, parent, parentPropName);
              break;
            case "null":
              if (val === null) {
                addType = true;
              }
              break;
            default:
              throw new TypeError("Unknown value type " + valueType);
          }
          if (addType) {
            retObj = {
              path,
              value: val,
              parent,
              parentProperty: parentPropName
            };
            this._handleCallback(retObj, callback, "value");
            return retObj;
          }
        } else if (loc[0] === "`" && val && hasOwnProp.call(val, loc.slice(1))) {
          var locProp = loc.slice(1);
          addRet(this._trace(x, val[locProp], push(path, locProp), val, locProp, callback, hasArrExpr, true));
        } else if (loc.includes(",")) {
          var parts = loc.split(",");
          var _iterator = _createForOfIteratorHelper(parts), _step;
          try {
            for (_iterator.s(); !(_step = _iterator.n()).done; ) {
              var part = _step.value;
              addRet(this._trace(unshift(part, x), val, path, parent, parentPropName, callback, true));
            }
          } catch (err) {
            _iterator.e(err);
          } finally {
            _iterator.f();
          }
        } else if (!literalPriority && val && hasOwnProp.call(val, loc)) {
          addRet(this._trace(x, val[loc], push(path, loc), val, loc, callback, hasArrExpr, true));
        }
        if (this._hasParentSelector) {
          for (var t = 0; t < ret.length; t++) {
            var rett = ret[t];
            if (rett && rett.isParentSelector) {
              var tmp = this._trace(rett.expr, val, rett.path, parent, parentPropName, callback, hasArrExpr);
              if (Array.isArray(tmp)) {
                ret[t] = tmp[0];
                var tl = tmp.length;
                for (var tt = 1; tt < tl; tt++) {
                  t++;
                  ret.splice(t, 0, tmp[tt]);
                }
              } else {
                ret[t] = tmp;
              }
            }
          }
        }
        return ret;
      };
      JSONPath.prototype._walk = function(loc, expr, val, path, parent, parentPropName, callback, f) {
        if (Array.isArray(val)) {
          var n = val.length;
          for (var i = 0; i < n; i++) {
            f(i, loc, expr, val, path, parent, parentPropName, callback);
          }
        } else if (val && _typeof(val) === "object") {
          Object.keys(val).forEach(function(m) {
            f(m, loc, expr, val, path, parent, parentPropName, callback);
          });
        }
      };
      JSONPath.prototype._slice = function(loc, expr, val, path, parent, parentPropName, callback) {
        if (!Array.isArray(val)) {
          return void 0;
        }
        var len = val.length, parts = loc.split(":"), step = parts[2] && Number.parseInt(parts[2]) || 1;
        var start = parts[0] && Number.parseInt(parts[0]) || 0, end = parts[1] && Number.parseInt(parts[1]) || len;
        start = start < 0 ? Math.max(0, start + len) : Math.min(len, start);
        end = end < 0 ? Math.max(0, end + len) : Math.min(len, end);
        var ret = [];
        for (var i = start; i < end; i += step) {
          var tmp = this._trace(unshift(i, expr), val, path, parent, parentPropName, callback, true);
          tmp.forEach(function(t) {
            ret.push(t);
          });
        }
        return ret;
      };
      JSONPath.prototype._eval = function(code, _v, _vname, path, parent, parentPropName) {
        if (code.includes("@parentProperty")) {
          this.currSandbox._$_parentProperty = parentPropName;
          code = code.replace(/@parentProperty/g, "_$_parentProperty");
        }
        if (code.includes("@parent")) {
          this.currSandbox._$_parent = parent;
          code = code.replace(/@parent/g, "_$_parent");
        }
        if (code.includes("@property")) {
          this.currSandbox._$_property = _vname;
          code = code.replace(/@property/g, "_$_property");
        }
        if (code.includes("@path")) {
          this.currSandbox._$_path = JSONPath.toPathString(path.concat([_vname]));
          code = code.replace(/@path/g, "_$_path");
        }
        if (code.includes("@root")) {
          this.currSandbox._$_root = this.json;
          code = code.replace(/@root/g, "_$_root");
        }
        if (/@([\t-\r \)\.\[\xA0\u1680\u2000-\u200A\u2028\u2029\u202F\u205F\u3000\uFEFF])/.test(code)) {
          this.currSandbox._$_v = _v;
          code = code.replace(/@([\t-\r \)\.\[\xA0\u1680\u2000-\u200A\u2028\u2029\u202F\u205F\u3000\uFEFF])/g, "_$_v$1");
        }
        try {
          return this.vm.runInNewContext(code, this.currSandbox);
        } catch (e) {
          console.log(e);
          throw new Error("jsonPath: " + e.message + ": " + code);
        }
      };
      JSONPath.cache = {};
      JSONPath.toPathString = function(pathArr) {
        var x = pathArr, n = x.length;
        var p = "$";
        for (var i = 1; i < n; i++) {
          if (!/^(~|\^|@(?:[\0-\t\x0B\f\x0E-\u2027\u202A-\uD7FF\uE000-\uFFFF]|[\uD800-\uDBFF][\uDC00-\uDFFF]|[\uD800-\uDBFF](?![\uDC00-\uDFFF])|(?:[^\uD800-\uDBFF]|^)[\uDC00-\uDFFF])*?\(\))$/.test(x[i])) {
            p += /^[\*0-9]+$/.test(x[i]) ? "[" + x[i] + "]" : "['" + x[i] + "']";
          }
        }
        return p;
      };
      JSONPath.toPointer = function(pointer) {
        var x = pointer, n = x.length;
        var p = "";
        for (var i = 1; i < n; i++) {
          if (!/^(~|\^|@(?:[\0-\t\x0B\f\x0E-\u2027\u202A-\uD7FF\uE000-\uFFFF]|[\uD800-\uDBFF][\uDC00-\uDFFF]|[\uD800-\uDBFF](?![\uDC00-\uDFFF])|(?:[^\uD800-\uDBFF]|^)[\uDC00-\uDFFF])*?\(\))$/.test(x[i])) {
            p += "/" + x[i].toString().replace(/~/g, "~0").replace(/\//g, "~1");
          }
        }
        return p;
      };
      JSONPath.toPathArray = function(expr) {
        var cache = JSONPath.cache;
        if (cache[expr]) {
          return cache[expr].concat();
        }
        var subx = [];
        var normalized = expr.replace(/@(?:null|boolean|number|string|integer|undefined|nonFinite|scalar|array|object|function|other)\(\)/g, ";$&;").replace(/['\[](\??\((?:[\0-\t\x0B\f\x0E-\u2027\u202A-\uD7FF\uE000-\uFFFF]|[\uD800-\uDBFF][\uDC00-\uDFFF]|[\uD800-\uDBFF](?![\uDC00-\uDFFF])|(?:[^\uD800-\uDBFF]|^)[\uDC00-\uDFFF])*?\))['\]]/g, function($0, $1) {
          return "[#" + (subx.push($1) - 1) + "]";
        }).replace(/\[["']((?:(?!['\]])[\s\S])*)["']\]/g, function($0, prop) {
          return "['" + prop.replace(/\./g, "%@%").replace(/~/g, "%%@@%%") + "']";
        }).replace(/~/g, ";~;").replace(/["']?\.["']?(?!(?:(?!\[)[\s\S])*\])|\[["']?/g, ";").replace(/%@%/g, ".").replace(/%%@@%%/g, "~").replace(/(?:;)?(\^+)(?:;)?/g, function($0, ups) {
          return ";" + ups.split("").join(";") + ";";
        }).replace(/;;;|;;/g, ";..;").replace(/;$|'?\]|'$/g, "");
        var exprList = normalized.split(";").map(function(exp) {
          var match = exp.match(/#([0-9]+)/);
          return !match || !match[1] ? exp : subx[match[1]];
        });
        cache[expr] = exprList;
        return cache[expr].concat();
      };
      moveToAnotherArray = function moveToAnotherArray2(source, target, conditionCb) {
        var il = source.length;
        for (var i = 0; i < il; i++) {
          var item = source[i];
          if (conditionCb(item)) {
            target.push(source.splice(i--, 1)[0]);
          }
        }
      };
      JSONPath.prototype.vm = {
        runInNewContext: function runInNewContext(expr, context) {
          var keys = Object.keys(context);
          var funcs = [];
          moveToAnotherArray(keys, funcs, function(key) {
            return typeof context[key] === "function";
          });
          var values = keys.map(function(vr, i) {
            return context[vr];
          });
          var funcString = funcs.reduce(function(s, func) {
            var fString = context[func].toString();
            if (!/function/.test(fString)) {
              fString = "function " + fString;
            }
            return "var " + func + "=" + fString + ";" + s;
          }, "");
          expr = funcString + expr;
          if (!/(["'])use strict\1/.test(expr) && !keys.includes("arguments")) {
            expr = "var arguments = undefined;" + expr;
          }
          expr = expr.replace(/;[\t-\r \xA0\u1680\u2000-\u200A\u2028\u2029\u202F\u205F\u3000\uFEFF]*$/, "");
          var lastStatementEnd = expr.lastIndexOf(";");
          var code = lastStatementEnd > -1 ? expr.slice(0, lastStatementEnd + 1) + " return " + expr.slice(lastStatementEnd + 1) : " return " + expr;
          return _construct(Function, _toConsumableArray(keys).concat([code])).apply(void 0, _toConsumableArray(values));
        }
      };
    }
  });

  // src/main.js
  var main_exports = {};
  __export(main_exports, {
    Client: () => Client,
    Response: () => Response,
    jsonpath: () => jsonpath
  });
  var jsonpath = (init_index_browser_esm(), index_browser_esm_exports);
  var Variables = class {
    constructor(s) {
      __publicField(this, "variableBag");
      this.variableBag = JSON.parse(s);
    }
    set(key, value) {
      this.variableBag[key] = value;
    }
    get(key) {
      return this.variableBag[key];
    }
  };
  var Client = class {
    constructor(session, environment, collection) {
      __publicField(this, "session");
      __publicField(this, "environment");
      __publicField(this, "collection");
      this.session = new Variables(session);
      this.environment = new Variables(environment);
      this.collection = new Variables(collection);
    }
  };
  var Response = class {
    constructor(body, headers, status) {
      __publicField(this, "body");
      __publicField(this, "headers");
      __publicField(this, "status");
      __publicField(this, "contentType");
      try {
        this.body = JSON.parse(body);
      } catch (error) {
        this.body = body;
      }
      this.headers = new Headers(headers);
      this.status = JSON.parse(status);
    }
  };
  var _data;
  var Headers = class {
    constructor(headers) {
      __privateAdd(this, _data, void 0);
      __privateSet(this, _data, JSON.parse(headers));
    }
    valueOf(s) {
      var _a;
      return (_a = __privateGet(this, _data)[s][0]) != null ? _a : null;
    }
    valuesOf(s) {
      var _a;
      return (_a = __privateGet(this, _data)[s]) != null ? _a : null;
    }
  };
  _data = new WeakMap();
  return main_exports;
})();
