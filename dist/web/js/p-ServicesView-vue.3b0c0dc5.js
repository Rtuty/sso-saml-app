"use strict";(self["webpackChunkandoma_passport_frontend"]=self["webpackChunkandoma_passport_frontend"]||[]).push([[565],{3326:function(e,t,s){s.r(t),s.d(t,{default:function(){return f}});var n=function(){var e=this,t=e._self._c;return t("b-container",{staticClass:"services",attrs:{fluid:""}},[t("h1",[e._v("Services")]),t("b-row",[t("b-col",{attrs:{sm:"2"}},[t("b-list-group",e._l(e.items,(function(s,n){return t("b-list-group-item",{key:`item-${n}`,attrs:{to:{name:"Services",query:{id:`${s}`}},exact:"","exact-active-class":"active"}},[e._v(" "+e._s(s)+" ")])})),1)],1),t("b-col",{attrs:{sm:"10"}},[t("pre",[e._v(e._s(e.info))])])],1)],1)},r=[],i=s(5102);const a=(e,t,s,n)=>{const r=t.childNodes;let i,o=null,c=1,l=null;for(let d=0;d<r.length;d++){if(l=r[d],i=l.nodeType,3!=i){if(3==c)o.nodeValue=n;else{const s=e.createTextNode(n);t.insertBefore(s,l),d++}if(1==i){const t=0==l.childNodes.length||1==l.childNodes.length&&1!=l.childNodes[0].nodeType;t||a(e,l,n,n+"  ")}}o=l,c=i}if(null!=l)if(3==i)l.nodeValue=s;else{const n=e.createTextNode(s);t.append(n)}},o=e=>{const t=(new DOMParser).parseFromString(e,"application/xml"),s=t.documentElement;return a(t,s,"\n","\n  "),(new XMLSerializer).serializeToString(t)};var c={name:"ServicesView",props:{id:{type:String}},data(){return{items:[],info:null}},created(){this.getItems()},watch:{id:{immediate:!0,handler(e){this.getInfo(e)}}},methods:{async getItems(){try{const e=await(0,i.x)().get("/services");if(e&&e.data){const{services:t}=e.data;Array.isArray(t)&&(this.items=t)}}catch(e){console.error(e)}},async getInfo(e){try{if(e){const t=await(0,i.x)().get(`/services/${e}`);t&&t.data&&(this.info=o(t.data))}else this.info=null}catch(t){console.error(t)}}}},l=c,d=s(1001),u=(0,d.Z)(l,n,r,!1,null,null,null),f=u.exports}}]);