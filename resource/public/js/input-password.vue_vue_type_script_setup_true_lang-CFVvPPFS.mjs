var C=Object.defineProperty;var f=Object.getOwnPropertySymbols;var b=Object.prototype.hasOwnProperty,V=Object.prototype.propertyIsEnumerable;var h=(s,t,e)=>t in s?C(s,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):s[t]=e,m=(s,t)=>{for(var e in t||(t={}))b.call(t,e)&&h(s,e,t[e]);if(f)for(var e of f(t))V.call(t,e)&&h(s,e,t[e]);return s};import{aF as o,aJ as l,aK as i,a6 as k,J as g,F as _,aY as B,bb as M,aZ as E,a_ as L,aB as S,P as z,n as v,aH as F,aa as u,aQ as $,aX as y,aU as I,aG as w}from"../jse/index-index-CmiOfTxA.mjs";import{c as x}from"./button.vue_vue_type_script_setup_true_lang-DJjG6SQm.mjs";import{I as D}from"./index-rzVbSPMH.mjs";import{_ as N}from"./input.vue_vue_type_script_setup_true_lang-dAgNxpnQ.mjs";const P=x("EyeOffIcon",[["path",{d:"M10.733 5.076a10.744 10.744 0 0 1 11.205 6.575 1 1 0 0 1 0 .696 10.747 10.747 0 0 1-1.444 2.49",key:"ct8e1f"}],["path",{d:"M14.084 14.158a3 3 0 0 1-4.242-4.242",key:"151rxh"}],["path",{d:"M17.479 17.499a10.75 10.75 0 0 1-15.417-5.151 1 1 0 0 1 0-.696 10.75 10.75 0 0 1 4.446-5.143",key:"13bj9a"}],["path",{d:"m2 2 20 20",key:"1ooewy"}]]);const T=x("EyeIcon",[["path",{d:"M2.062 12.348a1 1 0 0 1 0-.696 10.75 10.75 0 0 1 19.876 0 1 1 0 0 1 0 .696 10.75 10.75 0 0 1-19.876 0",key:"1nclc0"}],["circle",{cx:"12",cy:"12",r:"3",key:"1v7zrd"}]]);function Q(s,t){return o(),l("svg",{width:"15",height:"15",viewBox:"0 0 15 15",fill:"none",xmlns:"http://www.w3.org/2000/svg"},[i("path",{"fill-rule":"evenodd","clip-rule":"evenodd",d:"M11.4669 3.72684C11.7558 3.91574 11.8369 4.30308 11.648 4.59198L7.39799 11.092C7.29783 11.2452 7.13556 11.3467 6.95402 11.3699C6.77247 11.3931 6.58989 11.3355 6.45446 11.2124L3.70446 8.71241C3.44905 8.48022 3.43023 8.08494 3.66242 7.82953C3.89461 7.57412 4.28989 7.55529 4.5453 7.78749L6.75292 9.79441L10.6018 3.90792C10.7907 3.61902 11.178 3.53795 11.4669 3.72684Z",fill:"currentColor"})])}const Z={class:"relative mt-2 flex items-center justify-between"},j=k({__name:"password-strength",props:{password:{default:""}},setup(s){const t=s,e=["","#e74242","#ED6F6F","#EFBD47","#55D18780","#55D187"],c=g(()=>n(t.password)),p=g(()=>e[c.value]);function n(r){let a=0;return r.length>=8&&a++,/[a-z]/.test(r)&&a++,/[A-Z]/.test(r)&&a++,/\d/.test(r)&&a++,/[^\da-z]/i.test(r)&&a++,a}return(r,a)=>(o(),l("div",Z,[(o(),l(_,null,B(5,d=>i("div",{key:d,class:"dark:bg-input-background bg-heavy relative mr-1 h-1.5 w-1/5 rounded-sm last:mr-0"},[i("span",{style:M({backgroundColor:p.value,width:c.value>=d?"100%":""}),class:"absolute left-0 h-full w-0 rounded-sm transition-all duration-500"},null,4)])),64))]))}}),A={class:"relative"},J={key:0,class:"text-muted-foreground mt-1.5 text-xs"},X=k({inheritAttrs:!1,__name:"input-password",props:E({class:{},errorTip:{},label:{},name:{},passwordStrength:{type:Boolean},status:{}},{modelValue:{},modelModifiers:{}}),emits:["update:modelValue"],setup(s){const t=s,e=L(s,"modelValue"),c=S(),p=D(t),n=z(!1);return(r,a)=>(o(),l("div",A,[v(u(N),I({modelValue:e.value,"onUpdate:modelValue":a[0]||(a[0]=d=>e.value=d)},m(m({},u(p)),r.$attrs),{type:n.value?"text":"password"}),{default:F(()=>[r.passwordStrength?(o(),l(_,{key:0},[v(j,{password:e.value},null,8,["password"]),u(c).strengthText?(o(),l("p",J,[$(r.$slots,"strengthText")])):y("",!0)],64)):y("",!0)]),_:3},16,["modelValue","type"]),i("div",{class:"hover:text-foreground text-foreground/60 absolute inset-y-0 right-0 top-3 flex cursor-pointer pr-3 text-lg leading-5",onClick:a[1]||(a[1]=d=>n.value=!n.value)},[n.value?(o(),w(u(T),{key:0,class:"size-4"})):(o(),w(u(P),{key:1,class:"size-4"}))])]))}});export{X as _,Q as r};
