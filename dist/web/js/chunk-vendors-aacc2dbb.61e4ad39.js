"use strict";(self["webpackChunkandoma_passport_frontend"]=self["webpackChunkandoma_passport_frontend"]||[]).push([[99],{902:function(t,e,n){
/**!
 * @fileOverview Kickass library to create and place poppers near their reference elements.
 * @version 1.16.1
 * @license
 * Copyright (c) 2016 Federico Zivolo and contributors
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
var r="undefined"!==typeof window&&"undefined"!==typeof document&&"undefined"!==typeof navigator,o=function(){for(var t=["Edge","Trident","Firefox"],e=0;e<t.length;e+=1)if(r&&navigator.userAgent.indexOf(t[e])>=0)return 1;return 0}();function i(t){var e=!1;return function(){e||(e=!0,window.Promise.resolve().then((function(){e=!1,t()})))}}function s(t){var e=!1;return function(){e||(e=!0,setTimeout((function(){e=!1,t()}),o))}}var a=r&&window.Promise,f=a?i:s;function p(t){var e={};return t&&"[object Function]"===e.toString.call(t)}function l(t,e){if(1!==t.nodeType)return[];var n=t.ownerDocument.defaultView,r=n.getComputedStyle(t,null);return e?r[e]:r}function u(t){return"HTML"===t.nodeName?t:t.parentNode||t.host}function c(t){if(!t)return document.body;switch(t.nodeName){case"HTML":case"BODY":return t.ownerDocument.body;case"#document":return t.body}var e=l(t),n=e.overflow,r=e.overflowX,o=e.overflowY;return/(auto|scroll|overlay)/.test(n+o+r)?t:c(u(t))}function d(t){return t&&t.referenceNode?t.referenceNode:t}var h=r&&!(!window.MSInputMethodContext||!document.documentMode),m=r&&/MSIE 10/.test(navigator.userAgent);function g(t){return 11===t?h:10===t?m:h||m}function v(t){if(!t)return document.documentElement;var e=g(10)?document.body:null,n=t.offsetParent||null;while(n===e&&t.nextElementSibling)n=(t=t.nextElementSibling).offsetParent;var r=n&&n.nodeName;return r&&"BODY"!==r&&"HTML"!==r?-1!==["TH","TD","TABLE"].indexOf(n.nodeName)&&"static"===l(n,"position")?v(n):n:t?t.ownerDocument.documentElement:document.documentElement}function b(t){var e=t.nodeName;return"BODY"!==e&&("HTML"===e||v(t.firstElementChild)===t)}function y(t){return null!==t.parentNode?y(t.parentNode):t}function w(t,e){if(!t||!t.nodeType||!e||!e.nodeType)return document.documentElement;var n=t.compareDocumentPosition(e)&Node.DOCUMENT_POSITION_FOLLOWING,r=n?t:e,o=n?e:t,i=document.createRange();i.setStart(r,0),i.setEnd(o,0);var s=i.commonAncestorContainer;if(t!==s&&e!==s||r.contains(o))return b(s)?s:v(s);var a=y(t);return a.host?w(a.host,e):w(t,y(e).host)}function x(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"top",n="top"===e?"scrollTop":"scrollLeft",r=t.nodeName;if("BODY"===r||"HTML"===r){var o=t.ownerDocument.documentElement,i=t.ownerDocument.scrollingElement||o;return i[n]}return t[n]}function O(t,e){var n=arguments.length>2&&void 0!==arguments[2]&&arguments[2],r=x(e,"top"),o=x(e,"left"),i=n?-1:1;return t.top+=r*i,t.bottom+=r*i,t.left+=o*i,t.right+=o*i,t}function E(t,e){var n="x"===e?"Left":"Top",r="Left"===n?"Right":"Bottom";return parseFloat(t["border"+n+"Width"])+parseFloat(t["border"+r+"Width"])}function S(t,e,n,r){return Math.max(e["offset"+t],e["scroll"+t],n["client"+t],n["offset"+t],n["scroll"+t],g(10)?parseInt(n["offset"+t])+parseInt(r["margin"+("Height"===t?"Top":"Left")])+parseInt(r["margin"+("Height"===t?"Bottom":"Right")]):0)}function T(t){var e=t.body,n=t.documentElement,r=g(10)&&getComputedStyle(n);return{height:S("Height",e,n,r),width:S("Width",e,n,r)}}var L=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")},C=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}(),D=function(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t},M=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var n=arguments[e];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(t[r]=n[r])}return t};function N(t){return M({},t,{right:t.left+t.width,bottom:t.top+t.height})}function P(t){var e={};try{if(g(10)){e=t.getBoundingClientRect();var n=x(t,"top"),r=x(t,"left");e.top+=n,e.left+=r,e.bottom+=n,e.right+=r}else e=t.getBoundingClientRect()}catch(c){}var o={left:e.left,top:e.top,width:e.right-e.left,height:e.bottom-e.top},i="HTML"===t.nodeName?T(t.ownerDocument):{},s=i.width||t.clientWidth||o.width,a=i.height||t.clientHeight||o.height,f=t.offsetWidth-s,p=t.offsetHeight-a;if(f||p){var u=l(t);f-=E(u,"x"),p-=E(u,"y"),o.width-=f,o.height-=p}return N(o)}function k(t,e){var n=arguments.length>2&&void 0!==arguments[2]&&arguments[2],r=g(10),o="HTML"===e.nodeName,i=P(t),s=P(e),a=c(t),f=l(e),p=parseFloat(f.borderTopWidth),u=parseFloat(f.borderLeftWidth);n&&o&&(s.top=Math.max(s.top,0),s.left=Math.max(s.left,0));var d=N({top:i.top-s.top-p,left:i.left-s.left-u,width:i.width,height:i.height});if(d.marginTop=0,d.marginLeft=0,!r&&o){var h=parseFloat(f.marginTop),m=parseFloat(f.marginLeft);d.top-=p-h,d.bottom-=p-h,d.left-=u-m,d.right-=u-m,d.marginTop=h,d.marginLeft=m}return(r&&!n?e.contains(a):e===a&&"BODY"!==a.nodeName)&&(d=O(d,e)),d}function F(t){var e=arguments.length>1&&void 0!==arguments[1]&&arguments[1],n=t.ownerDocument.documentElement,r=k(t,n),o=Math.max(n.clientWidth,window.innerWidth||0),i=Math.max(n.clientHeight,window.innerHeight||0),s=e?0:x(n),a=e?0:x(n,"left"),f={top:s-r.top+r.marginTop,left:a-r.left+r.marginLeft,width:o,height:i};return N(f)}function A(t){var e=t.nodeName;if("BODY"===e||"HTML"===e)return!1;if("fixed"===l(t,"position"))return!0;var n=u(t);return!!n&&A(n)}function B(t){if(!t||!t.parentElement||g())return document.documentElement;var e=t.parentElement;while(e&&"none"===l(e,"transform"))e=e.parentElement;return e||document.documentElement}function $(t,e,n,r){var o=arguments.length>4&&void 0!==arguments[4]&&arguments[4],i={top:0,left:0},s=o?B(t):w(t,d(e));if("viewport"===r)i=F(s,o);else{var a=void 0;"scrollParent"===r?(a=c(u(e)),"BODY"===a.nodeName&&(a=t.ownerDocument.documentElement)):a="window"===r?t.ownerDocument.documentElement:r;var f=k(a,s,o);if("HTML"!==a.nodeName||A(s))i=f;else{var p=T(t.ownerDocument),l=p.height,h=p.width;i.top+=f.top-f.marginTop,i.bottom=l+f.top,i.left+=f.left-f.marginLeft,i.right=h+f.left}}n=n||0;var m="number"===typeof n;return i.left+=m?n:n.left||0,i.top+=m?n:n.top||0,i.right-=m?n:n.right||0,i.bottom-=m?n:n.bottom||0,i}function j(t){var e=t.width,n=t.height;return e*n}function I(t,e,n,r,o){var i=arguments.length>5&&void 0!==arguments[5]?arguments[5]:0;if(-1===t.indexOf("auto"))return t;var s=$(n,r,i,o),a={top:{width:s.width,height:e.top-s.top},right:{width:s.right-e.right,height:s.height},bottom:{width:s.width,height:s.bottom-e.bottom},left:{width:e.left-s.left,height:s.height}},f=Object.keys(a).map((function(t){return M({key:t},a[t],{area:j(a[t])})})).sort((function(t,e){return e.area-t.area})),p=f.filter((function(t){var e=t.width,r=t.height;return e>=n.clientWidth&&r>=n.clientHeight})),l=p.length>0?p[0].key:f[0].key,u=t.split("-")[1];return l+(u?"-"+u:"")}function W(t,e,n){var r=arguments.length>3&&void 0!==arguments[3]?arguments[3]:null,o=r?B(e):w(e,d(n));return k(n,o,r)}function H(t){var e=t.ownerDocument.defaultView,n=e.getComputedStyle(t),r=parseFloat(n.marginTop||0)+parseFloat(n.marginBottom||0),o=parseFloat(n.marginLeft||0)+parseFloat(n.marginRight||0),i={width:t.offsetWidth+o,height:t.offsetHeight+r};return i}function R(t){var e={left:"right",right:"left",bottom:"top",top:"bottom"};return t.replace(/left|right|bottom|top/g,(function(t){return e[t]}))}function U(t,e,n){n=n.split("-")[0];var r=H(t),o={width:r.width,height:r.height},i=-1!==["right","left"].indexOf(n),s=i?"top":"left",a=i?"left":"top",f=i?"height":"width",p=i?"width":"height";return o[s]=e[s]+e[f]/2-r[f]/2,o[a]=n===a?e[a]-r[p]:e[R(a)],o}function z(t,e){return Array.prototype.find?t.find(e):t.filter(e)[0]}function V(t,e,n){if(Array.prototype.findIndex)return t.findIndex((function(t){return t[e]===n}));var r=z(t,(function(t){return t[e]===n}));return t.indexOf(r)}function _(t,e,n){var r=void 0===n?t:t.slice(0,V(t,"name",n));return r.forEach((function(t){t["function"]&&console.warn("`modifier.function` is deprecated, use `modifier.fn`!");var n=t["function"]||t.fn;t.enabled&&p(n)&&(e.offsets.popper=N(e.offsets.popper),e.offsets.reference=N(e.offsets.reference),e=n(e,t))})),e}function Y(){if(!this.state.isDestroyed){var t={instance:this,styles:{},arrowStyles:{},attributes:{},flipped:!1,offsets:{}};t.offsets.reference=W(this.state,this.popper,this.reference,this.options.positionFixed),t.placement=I(this.options.placement,t.offsets.reference,this.popper,this.reference,this.options.modifiers.flip.boundariesElement,this.options.modifiers.flip.padding),t.originalPlacement=t.placement,t.positionFixed=this.options.positionFixed,t.offsets.popper=U(this.popper,t.offsets.reference,t.placement),t.offsets.popper.position=this.options.positionFixed?"fixed":"absolute",t=_(this.modifiers,t),this.state.isCreated?this.options.onUpdate(t):(this.state.isCreated=!0,this.options.onCreate(t))}}function q(t,e){return t.some((function(t){var n=t.name,r=t.enabled;return r&&n===e}))}function K(t){for(var e=[!1,"ms","Webkit","Moz","O"],n=t.charAt(0).toUpperCase()+t.slice(1),r=0;r<e.length;r++){var o=e[r],i=o?""+o+n:t;if("undefined"!==typeof document.body.style[i])return i}return null}function G(){return this.state.isDestroyed=!0,q(this.modifiers,"applyStyle")&&(this.popper.removeAttribute("x-placement"),this.popper.style.position="",this.popper.style.top="",this.popper.style.left="",this.popper.style.right="",this.popper.style.bottom="",this.popper.style.willChange="",this.popper.style[K("transform")]=""),this.disableEventListeners(),this.options.removeOnDestroy&&this.popper.parentNode.removeChild(this.popper),this}function X(t){var e=t.ownerDocument;return e?e.defaultView:window}function Z(t,e,n,r){var o="BODY"===t.nodeName,i=o?t.ownerDocument.defaultView:t;i.addEventListener(e,n,{passive:!0}),o||Z(c(i.parentNode),e,n,r),r.push(i)}function J(t,e,n,r){n.updateBound=r,X(t).addEventListener("resize",n.updateBound,{passive:!0});var o=c(t);return Z(o,"scroll",n.updateBound,n.scrollParents),n.scrollElement=o,n.eventsEnabled=!0,n}function Q(){this.state.eventsEnabled||(this.state=J(this.reference,this.options,this.state,this.scheduleUpdate))}function tt(t,e){return X(t).removeEventListener("resize",e.updateBound),e.scrollParents.forEach((function(t){t.removeEventListener("scroll",e.updateBound)})),e.updateBound=null,e.scrollParents=[],e.scrollElement=null,e.eventsEnabled=!1,e}function et(){this.state.eventsEnabled&&(cancelAnimationFrame(this.scheduleUpdate),this.state=tt(this.reference,this.state))}function nt(t){return""!==t&&!isNaN(parseFloat(t))&&isFinite(t)}function rt(t,e){Object.keys(e).forEach((function(n){var r="";-1!==["width","height","top","right","bottom","left"].indexOf(n)&&nt(e[n])&&(r="px"),t.style[n]=e[n]+r}))}function ot(t,e){Object.keys(e).forEach((function(n){var r=e[n];!1!==r?t.setAttribute(n,e[n]):t.removeAttribute(n)}))}function it(t){return rt(t.instance.popper,t.styles),ot(t.instance.popper,t.attributes),t.arrowElement&&Object.keys(t.arrowStyles).length&&rt(t.arrowElement,t.arrowStyles),t}function st(t,e,n,r,o){var i=W(o,e,t,n.positionFixed),s=I(n.placement,i,e,t,n.modifiers.flip.boundariesElement,n.modifiers.flip.padding);return e.setAttribute("x-placement",s),rt(e,{position:n.positionFixed?"fixed":"absolute"}),n}function at(t,e){var n=t.offsets,r=n.popper,o=n.reference,i=Math.round,s=Math.floor,a=function(t){return t},f=i(o.width),p=i(r.width),l=-1!==["left","right"].indexOf(t.placement),u=-1!==t.placement.indexOf("-"),c=f%2===p%2,d=f%2===1&&p%2===1,h=e?l||u||c?i:s:a,m=e?i:a;return{left:h(d&&!u&&e?r.left-1:r.left),top:m(r.top),bottom:m(r.bottom),right:h(r.right)}}var ft=r&&/Firefox/i.test(navigator.userAgent);function pt(t,e){var n=e.x,r=e.y,o=t.offsets.popper,i=z(t.instance.modifiers,(function(t){return"applyStyle"===t.name})).gpuAcceleration;void 0!==i&&console.warn("WARNING: `gpuAcceleration` option moved to `computeStyle` modifier and will not be supported in future versions of Popper.js!");var s=void 0!==i?i:e.gpuAcceleration,a=v(t.instance.popper),f=P(a),p={position:o.position},l=at(t,window.devicePixelRatio<2||!ft),u="bottom"===n?"top":"bottom",c="right"===r?"left":"right",d=K("transform"),h=void 0,m=void 0;if(m="bottom"===u?"HTML"===a.nodeName?-a.clientHeight+l.bottom:-f.height+l.bottom:l.top,h="right"===c?"HTML"===a.nodeName?-a.clientWidth+l.right:-f.width+l.right:l.left,s&&d)p[d]="translate3d("+h+"px, "+m+"px, 0)",p[u]=0,p[c]=0,p.willChange="transform";else{var g="bottom"===u?-1:1,b="right"===c?-1:1;p[u]=m*g,p[c]=h*b,p.willChange=u+", "+c}var y={"x-placement":t.placement};return t.attributes=M({},y,t.attributes),t.styles=M({},p,t.styles),t.arrowStyles=M({},t.offsets.arrow,t.arrowStyles),t}function lt(t,e,n){var r=z(t,(function(t){var n=t.name;return n===e})),o=!!r&&t.some((function(t){return t.name===n&&t.enabled&&t.order<r.order}));if(!o){var i="`"+e+"`",s="`"+n+"`";console.warn(s+" modifier is required by "+i+" modifier in order to work, be sure to include it before "+i+"!")}return o}function ut(t,e){var n;if(!lt(t.instance.modifiers,"arrow","keepTogether"))return t;var r=e.element;if("string"===typeof r){if(r=t.instance.popper.querySelector(r),!r)return t}else if(!t.instance.popper.contains(r))return console.warn("WARNING: `arrow.element` must be child of its popper element!"),t;var o=t.placement.split("-")[0],i=t.offsets,s=i.popper,a=i.reference,f=-1!==["left","right"].indexOf(o),p=f?"height":"width",u=f?"Top":"Left",c=u.toLowerCase(),d=f?"left":"top",h=f?"bottom":"right",m=H(r)[p];a[h]-m<s[c]&&(t.offsets.popper[c]-=s[c]-(a[h]-m)),a[c]+m>s[h]&&(t.offsets.popper[c]+=a[c]+m-s[h]),t.offsets.popper=N(t.offsets.popper);var g=a[c]+a[p]/2-m/2,v=l(t.instance.popper),b=parseFloat(v["margin"+u]),y=parseFloat(v["border"+u+"Width"]),w=g-t.offsets.popper[c]-b-y;return w=Math.max(Math.min(s[p]-m,w),0),t.arrowElement=r,t.offsets.arrow=(n={},D(n,c,Math.round(w)),D(n,d,""),n),t}function ct(t){return"end"===t?"start":"start"===t?"end":t}var dt=["auto-start","auto","auto-end","top-start","top","top-end","right-start","right","right-end","bottom-end","bottom","bottom-start","left-end","left","left-start"],ht=dt.slice(3);function mt(t){var e=arguments.length>1&&void 0!==arguments[1]&&arguments[1],n=ht.indexOf(t),r=ht.slice(n+1).concat(ht.slice(0,n));return e?r.reverse():r}var gt={FLIP:"flip",CLOCKWISE:"clockwise",COUNTERCLOCKWISE:"counterclockwise"};function vt(t,e){if(q(t.instance.modifiers,"inner"))return t;if(t.flipped&&t.placement===t.originalPlacement)return t;var n=$(t.instance.popper,t.instance.reference,e.padding,e.boundariesElement,t.positionFixed),r=t.placement.split("-")[0],o=R(r),i=t.placement.split("-")[1]||"",s=[];switch(e.behavior){case gt.FLIP:s=[r,o];break;case gt.CLOCKWISE:s=mt(r);break;case gt.COUNTERCLOCKWISE:s=mt(r,!0);break;default:s=e.behavior}return s.forEach((function(a,f){if(r!==a||s.length===f+1)return t;r=t.placement.split("-")[0],o=R(r);var p=t.offsets.popper,l=t.offsets.reference,u=Math.floor,c="left"===r&&u(p.right)>u(l.left)||"right"===r&&u(p.left)<u(l.right)||"top"===r&&u(p.bottom)>u(l.top)||"bottom"===r&&u(p.top)<u(l.bottom),d=u(p.left)<u(n.left),h=u(p.right)>u(n.right),m=u(p.top)<u(n.top),g=u(p.bottom)>u(n.bottom),v="left"===r&&d||"right"===r&&h||"top"===r&&m||"bottom"===r&&g,b=-1!==["top","bottom"].indexOf(r),y=!!e.flipVariations&&(b&&"start"===i&&d||b&&"end"===i&&h||!b&&"start"===i&&m||!b&&"end"===i&&g),w=!!e.flipVariationsByContent&&(b&&"start"===i&&h||b&&"end"===i&&d||!b&&"start"===i&&g||!b&&"end"===i&&m),x=y||w;(c||v||x)&&(t.flipped=!0,(c||v)&&(r=s[f+1]),x&&(i=ct(i)),t.placement=r+(i?"-"+i:""),t.offsets.popper=M({},t.offsets.popper,U(t.instance.popper,t.offsets.reference,t.placement)),t=_(t.instance.modifiers,t,"flip"))})),t}function bt(t){var e=t.offsets,n=e.popper,r=e.reference,o=t.placement.split("-")[0],i=Math.floor,s=-1!==["top","bottom"].indexOf(o),a=s?"right":"bottom",f=s?"left":"top",p=s?"width":"height";return n[a]<i(r[f])&&(t.offsets.popper[f]=i(r[f])-n[p]),n[f]>i(r[a])&&(t.offsets.popper[f]=i(r[a])),t}function yt(t,e,n,r){var o=t.match(/((?:\-|\+)?\d*\.?\d*)(.*)/),i=+o[1],s=o[2];if(!i)return t;if(0===s.indexOf("%")){var a=void 0;switch(s){case"%p":a=n;break;case"%":case"%r":default:a=r}var f=N(a);return f[e]/100*i}if("vh"===s||"vw"===s){var p=void 0;return p="vh"===s?Math.max(document.documentElement.clientHeight,window.innerHeight||0):Math.max(document.documentElement.clientWidth,window.innerWidth||0),p/100*i}return i}function wt(t,e,n,r){var o=[0,0],i=-1!==["right","left"].indexOf(r),s=t.split(/(\+|\-)/).map((function(t){return t.trim()})),a=s.indexOf(z(s,(function(t){return-1!==t.search(/,|\s/)})));s[a]&&-1===s[a].indexOf(",")&&console.warn("Offsets separated by white space(s) are deprecated, use a comma (,) instead.");var f=/\s*,\s*|\s+/,p=-1!==a?[s.slice(0,a).concat([s[a].split(f)[0]]),[s[a].split(f)[1]].concat(s.slice(a+1))]:[s];return p=p.map((function(t,r){var o=(1===r?!i:i)?"height":"width",s=!1;return t.reduce((function(t,e){return""===t[t.length-1]&&-1!==["+","-"].indexOf(e)?(t[t.length-1]=e,s=!0,t):s?(t[t.length-1]+=e,s=!1,t):t.concat(e)}),[]).map((function(t){return yt(t,o,e,n)}))})),p.forEach((function(t,e){t.forEach((function(n,r){nt(n)&&(o[e]+=n*("-"===t[r-1]?-1:1))}))})),o}function xt(t,e){var n=e.offset,r=t.placement,o=t.offsets,i=o.popper,s=o.reference,a=r.split("-")[0],f=void 0;return f=nt(+n)?[+n,0]:wt(n,i,s,a),"left"===a?(i.top+=f[0],i.left-=f[1]):"right"===a?(i.top+=f[0],i.left+=f[1]):"top"===a?(i.left+=f[0],i.top-=f[1]):"bottom"===a&&(i.left+=f[0],i.top+=f[1]),t.popper=i,t}function Ot(t,e){var n=e.boundariesElement||v(t.instance.popper);t.instance.reference===n&&(n=v(n));var r=K("transform"),o=t.instance.popper.style,i=o.top,s=o.left,a=o[r];o.top="",o.left="",o[r]="";var f=$(t.instance.popper,t.instance.reference,e.padding,n,t.positionFixed);o.top=i,o.left=s,o[r]=a,e.boundaries=f;var p=e.priority,l=t.offsets.popper,u={primary:function(t){var n=l[t];return l[t]<f[t]&&!e.escapeWithReference&&(n=Math.max(l[t],f[t])),D({},t,n)},secondary:function(t){var n="right"===t?"left":"top",r=l[n];return l[t]>f[t]&&!e.escapeWithReference&&(r=Math.min(l[n],f[t]-("right"===t?l.width:l.height))),D({},n,r)}};return p.forEach((function(t){var e=-1!==["left","top"].indexOf(t)?"primary":"secondary";l=M({},l,u[e](t))})),t.offsets.popper=l,t}function Et(t){var e=t.placement,n=e.split("-")[0],r=e.split("-")[1];if(r){var o=t.offsets,i=o.reference,s=o.popper,a=-1!==["bottom","top"].indexOf(n),f=a?"left":"top",p=a?"width":"height",l={start:D({},f,i[f]),end:D({},f,i[f]+i[p]-s[p])};t.offsets.popper=M({},s,l[r])}return t}function St(t){if(!lt(t.instance.modifiers,"hide","preventOverflow"))return t;var e=t.offsets.reference,n=z(t.instance.modifiers,(function(t){return"preventOverflow"===t.name})).boundaries;if(e.bottom<n.top||e.left>n.right||e.top>n.bottom||e.right<n.left){if(!0===t.hide)return t;t.hide=!0,t.attributes["x-out-of-boundaries"]=""}else{if(!1===t.hide)return t;t.hide=!1,t.attributes["x-out-of-boundaries"]=!1}return t}function Tt(t){var e=t.placement,n=e.split("-")[0],r=t.offsets,o=r.popper,i=r.reference,s=-1!==["left","right"].indexOf(n),a=-1===["top","left"].indexOf(n);return o[s?"left":"top"]=i[n]-(a?o[s?"width":"height"]:0),t.placement=R(e),t.offsets.popper=N(o),t}var Lt={shift:{order:100,enabled:!0,fn:Et},offset:{order:200,enabled:!0,fn:xt,offset:0},preventOverflow:{order:300,enabled:!0,fn:Ot,priority:["left","right","top","bottom"],padding:5,boundariesElement:"scrollParent"},keepTogether:{order:400,enabled:!0,fn:bt},arrow:{order:500,enabled:!0,fn:ut,element:"[x-arrow]"},flip:{order:600,enabled:!0,fn:vt,behavior:"flip",padding:5,boundariesElement:"viewport",flipVariations:!1,flipVariationsByContent:!1},inner:{order:700,enabled:!1,fn:Tt},hide:{order:800,enabled:!0,fn:St},computeStyle:{order:850,enabled:!0,fn:pt,gpuAcceleration:!0,x:"bottom",y:"right"},applyStyle:{order:900,enabled:!0,fn:it,onLoad:st,gpuAcceleration:void 0}},Ct={placement:"bottom",positionFixed:!1,eventsEnabled:!0,removeOnDestroy:!1,onCreate:function(){},onUpdate:function(){},modifiers:Lt},Dt=function(){function t(e,n){var r=this,o=arguments.length>2&&void 0!==arguments[2]?arguments[2]:{};L(this,t),this.scheduleUpdate=function(){return requestAnimationFrame(r.update)},this.update=f(this.update.bind(this)),this.options=M({},t.Defaults,o),this.state={isDestroyed:!1,isCreated:!1,scrollParents:[]},this.reference=e&&e.jquery?e[0]:e,this.popper=n&&n.jquery?n[0]:n,this.options.modifiers={},Object.keys(M({},t.Defaults.modifiers,o.modifiers)).forEach((function(e){r.options.modifiers[e]=M({},t.Defaults.modifiers[e]||{},o.modifiers?o.modifiers[e]:{})})),this.modifiers=Object.keys(this.options.modifiers).map((function(t){return M({name:t},r.options.modifiers[t])})).sort((function(t,e){return t.order-e.order})),this.modifiers.forEach((function(t){t.enabled&&p(t.onLoad)&&t.onLoad(r.reference,r.popper,r.options,t,r.state)})),this.update();var i=this.options.eventsEnabled;i&&this.enableEventListeners(),this.state.eventsEnabled=i}return C(t,[{key:"update",value:function(){return Y.call(this)}},{key:"destroy",value:function(){return G.call(this)}},{key:"enableEventListeners",value:function(){return Q.call(this)}},{key:"disableEventListeners",value:function(){return et.call(this)}}]),t}();Dt.Utils=("undefined"!==typeof window?window:n.g).PopperUtils,Dt.placements=dt,Dt.Defaults=Ct,e["Z"]=Dt},9342:function(t,e,n){function r(t){return t&&"object"===typeof t&&"default"in t?t["default"]:t}var o=r(n(7195));function i(t){return i="function"===typeof Symbol&&"symbol"===typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"===typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t},i(t)}function s(t){return a(t)||f(t)||p()}function a(t){if(Array.isArray(t)){for(var e=0,n=new Array(t.length);e<t.length;e++)n[e]=t[e];return n}}function f(t){if(Symbol.iterator in Object(t)||"[object Arguments]"===Object.prototype.toString.call(t))return Array.from(t)}function p(){throw new TypeError("Invalid attempt to spread non-iterable instance")}var l="undefined"!==typeof window;function u(t){return Array.isArray(t)||"object"===i(t)?Object.freeze(t):t}function c(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return t.reduce((function(t,n){var r=n.passengers[0],o="function"===typeof r?r(e):n.passengers;return t.concat(o)}),[])}function d(t,e){return t.map((function(t,e){return[e,t]})).sort((function(t,n){return e(t[1],n[1])||t[0]-n[0]})).map((function(t){return t[1]}))}function h(t,e){return e.reduce((function(e,n){return t.hasOwnProperty(n)&&(e[n]=t[n]),e}),{})}var m={},g={},v={},b=o.extend({data:function(){return{transports:m,targets:g,sources:v,trackInstances:l}},methods:{open:function(t){if(l){var e=t.to,n=t.from,r=t.passengers,i=t.order,s=void 0===i?1/0:i;if(e&&n&&r){var a={to:e,from:n,passengers:u(r),order:s},f=Object.keys(this.transports);-1===f.indexOf(e)&&o.set(this.transports,e,[]);var p=this.$_getTransportIndex(a),c=this.transports[e].slice(0);-1===p?c.push(a):c[p]=a,this.transports[e]=d(c,(function(t,e){return t.order-e.order}))}}},close:function(t){var e=arguments.length>1&&void 0!==arguments[1]&&arguments[1],n=t.to,r=t.from;if(n&&(r||!1!==e)&&this.transports[n])if(e)this.transports[n]=[];else{var o=this.$_getTransportIndex(t);if(o>=0){var i=this.transports[n].slice(0);i.splice(o,1),this.transports[n]=i}}},registerTarget:function(t,e,n){l&&(this.trackInstances&&!n&&this.targets[t]&&console.warn("[portal-vue]: Target ".concat(t," already exists")),this.$set(this.targets,t,Object.freeze([e])))},unregisterTarget:function(t){this.$delete(this.targets,t)},registerSource:function(t,e,n){l&&(this.trackInstances&&!n&&this.sources[t]&&console.warn("[portal-vue]: source ".concat(t," already exists")),this.$set(this.sources,t,Object.freeze([e])))},unregisterSource:function(t){this.$delete(this.sources,t)},hasTarget:function(t){return!(!this.targets[t]||!this.targets[t][0])},hasSource:function(t){return!(!this.sources[t]||!this.sources[t][0])},hasContentFor:function(t){return!!this.transports[t]&&!!this.transports[t].length},$_getTransportIndex:function(t){var e=t.to,n=t.from;for(var r in this.transports[e])if(this.transports[e][r].from===n)return+r;return-1}}}),y=new b(m),w=1,x=o.extend({name:"portal",props:{disabled:{type:Boolean},name:{type:String,default:function(){return String(w++)}},order:{type:Number,default:0},slim:{type:Boolean},slotProps:{type:Object,default:function(){return{}}},tag:{type:String,default:"DIV"},to:{type:String,default:function(){return String(Math.round(1e7*Math.random()))}}},created:function(){var t=this;this.$nextTick((function(){y.registerSource(t.name,t)}))},mounted:function(){this.disabled||this.sendUpdate()},updated:function(){this.disabled?this.clear():this.sendUpdate()},beforeDestroy:function(){y.unregisterSource(this.name),this.clear()},watch:{to:function(t,e){e&&e!==t&&this.clear(e),this.sendUpdate()}},methods:{clear:function(t){var e={from:this.name,to:t||this.to};y.close(e)},normalizeSlots:function(){return this.$scopedSlots.default?[this.$scopedSlots.default]:this.$slots.default},normalizeOwnChildren:function(t){return"function"===typeof t?t(this.slotProps):t},sendUpdate:function(){var t=this.normalizeSlots();if(t){var e={from:this.name,to:this.to,passengers:s(t),order:this.order};y.open(e)}else this.clear()}},render:function(t){var e=this.$slots.default||this.$scopedSlots.default||[],n=this.tag;return e&&this.disabled?e.length<=1&&this.slim?this.normalizeOwnChildren(e)[0]:t(n,[this.normalizeOwnChildren(e)]):this.slim?t():t(n,{class:{"v-portal":!0},style:{display:"none"},key:"v-portal-placeholder"})}}),O=o.extend({name:"portalTarget",props:{multiple:{type:Boolean,default:!1},name:{type:String,required:!0},slim:{type:Boolean,default:!1},slotProps:{type:Object,default:function(){return{}}},tag:{type:String,default:"div"},transition:{type:[String,Object,Function]}},data:function(){return{transports:y.transports,firstRender:!0}},created:function(){var t=this;this.$nextTick((function(){y.registerTarget(t.name,t)}))},watch:{ownTransports:function(){this.$emit("change",this.children().length>0)},name:function(t,e){y.unregisterTarget(e),y.registerTarget(t,this)}},mounted:function(){var t=this;this.transition&&this.$nextTick((function(){t.firstRender=!1}))},beforeDestroy:function(){y.unregisterTarget(this.name)},computed:{ownTransports:function(){var t=this.transports[this.name]||[];return this.multiple?t:0===t.length?[]:[t[t.length-1]]},passengers:function(){return c(this.ownTransports,this.slotProps)}},methods:{children:function(){return 0!==this.passengers.length?this.passengers:this.$scopedSlots.default?this.$scopedSlots.default(this.slotProps):this.$slots.default||[]},noWrapper:function(){var t=this.slim&&!this.transition;return t&&this.children().length>1&&console.warn("[portal-vue]: PortalTarget with `slim` option received more than one child element."),t}},render:function(t){var e=this.noWrapper(),n=this.children(),r=this.transition||this.tag;return e?n[0]:this.slim&&!r?t():t(r,{props:{tag:this.transition&&this.tag?this.tag:void 0},class:{"vue-portal-target":!0}},n)}}),E=0,S=["disabled","name","order","slim","slotProps","tag","to"],T=["multiple","transition"],L=o.extend({name:"MountingPortal",inheritAttrs:!1,props:{append:{type:[Boolean,String]},bail:{type:Boolean},mountTo:{type:String,required:!0},disabled:{type:Boolean},name:{type:String,default:function(){return"mounted_"+String(E++)}},order:{type:Number,default:0},slim:{type:Boolean},slotProps:{type:Object,default:function(){return{}}},tag:{type:String,default:"DIV"},to:{type:String,default:function(){return String(Math.round(1e7*Math.random()))}},multiple:{type:Boolean,default:!1},targetSlim:{type:Boolean},targetSlotProps:{type:Object,default:function(){return{}}},targetTag:{type:String,default:"div"},transition:{type:[String,Object,Function]}},created:function(){if("undefined"!==typeof document){var t=document.querySelector(this.mountTo);if(t){var e=this.$props;if(y.targets[e.name])e.bail?console.warn("[portal-vue]: Target ".concat(e.name," is already mounted.\n        Aborting because 'bail: true' is set")):this.portalTarget=y.targets[e.name];else{var n=e.append;if(n){var r="string"===typeof n?n:"DIV",o=document.createElement(r);t.appendChild(o),t=o}var i=h(this.$props,T);i.slim=this.targetSlim,i.tag=this.targetTag,i.slotProps=this.targetSlotProps,i.name=this.to,this.portalTarget=new O({el:t,parent:this.$parent||this,propsData:i})}}else console.error("[portal-vue]: Mount Point '".concat(this.mountTo,"' not found in document"))}},beforeDestroy:function(){var t=this.portalTarget;if(this.append){var e=t.$el;e.parentNode.removeChild(e)}t.$destroy()},render:function(t){if(!this.portalTarget)return console.warn("[portal-vue] Target wasn't mounted"),t();if(!this.$scopedSlots.manual){var e=h(this.$props,S);return t(x,{props:e,attrs:this.$attrs,on:this.$listeners,scopedSlots:this.$scopedSlots},this.$slots.default)}var n=this.$scopedSlots.manual({to:this.to});return Array.isArray(n)&&(n=n[0]),n||t()}});function C(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};t.component(e.portalName||"Portal",x),t.component(e.portalTargetName||"PortalTarget",O),t.component(e.MountingPortalName||"MountingPortal",L)}var D={install:C};e.h_=x,e.YC=O,e.Df=y}}]);