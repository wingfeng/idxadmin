var w=($,O,i)=>new Promise((c,t)=>{var C=n=>{try{p(i.next(n))}catch(_){t(_)}},D=n=>{try{p(i.throw(n))}catch(_){t(_)}},p=n=>n.done?c(n.value):Promise.resolve(n.value).then(C,D);p((i=i.apply($,O)).next())});import{u as Z}from"./index.es-CT2W4LN0.mjs";import{g as ee,d as ae,s as oe}from"./UserTable.vue_vue_type_script_setup_true_lang-DFfsm67F.mjs";import{n as te}from"./util-oQdYwX59.mjs";import{S as x}from"./enum-BLvisX2J.mjs";import{_ as le}from"./RoleForm.vue_vue_type_script_setup_true_lang-D5BtTssQ.mjs";import{S as ne,D as I,E as re,e as ie,M as se}from"./bootstrap-5eopxyZm.mjs";import{R as ue,M as de}from"./RollbackOutlined-C4ajMdEa.mjs";import{a6 as ce,P as d,J as M,a7 as pe,aI as s,aF as A,aJ as me,n as o,aH as l,aa as r,ao as k,h as b,aG as ve,aX as fe}from"../jse/index-index-CmiOfTxA.mjs";import"./user-7A8qe_c7.mjs";const _e={class:"p-5"},Re=ce({__name:"RoleList",setup($){const O=[{title:"Id",dataIndex:"id",key:"Id",sorter:!0,defaultSortOrder:"ascend"},{title:"Name",dataIndex:"name",key:"Name",name:"Name",sorter:!0,filtered:!0},{title:"Description",dataIndex:"description",key:"Description",filtered:!0},{title:"Actions",key:"action"}],i=d("id"),c=d(x.Asc),t=d({Name:"",Description:""}),C=M(()=>{const e=[];return t.value.Name!==""&&e.push("name like ?"),t.value.Description!==""&&e.push("description like ?"),e}),D=M(()=>{const e=[];return t.value.Name!==""&&e.push(`%${t.value.Name}%`),t.value.Description!==""&&e.push(`%${t.value.Description}%`),e}),{data:p,run:n,loading:_,total:T,current:h,pageSize:N}=Z(ee,{defaultParams:[{page:1,pageSize:10,sortField:"id",sortOrder:x.Asc,filters:C.value,args:D.value}],pagination:{currentKey:"page",pageSizeKey:"pageSize"}}),B=M(()=>({total:T.value,current:h.value,pageSize:N.value})),E=d(),g=()=>{setTimeout(()=>{n({page:h.value,pageSize:N.value,sortField:i.value,sortOrder:c.value,filters:C.value,args:D.value})},500)},P=()=>{console.log("searchModel",t.value),g()},U=()=>{t.value={Name:"",Description:""},E.value.resetFields(),g()},m=d(!1),S=d(),R=d(),V=e=>{S.value=e,m.value=!0},J=(e,a)=>{se.confirm({title:`Deleting Role ${e}`,content:`Are you sure delete this role ${a}?`,onOk(){ae(e),g()},onCancel(){console.log("Cancel")}}),console.log("record",e)},K=()=>w(this,null,function*(){try{const e=yield R.value.validate();if(!e)return;console.log("form validate",e)}catch(e){console.log("error submit:",e);return}oe(S.value),m.value=!1,g()}),L=()=>{R.value.resetForm(),m.value=!1},G=()=>w(this,null,function*(){const e=yield te();S.value={id:e.id,name:"",description:""},m.value=!0}),H=(e,a,v)=>{if(h.value=e.current,N.value=e.pageSize,v.field){const f=v.field;i.value=f!=null?f:v.field,c.value=v.order}else i.value="id",c.value=x.Asc;console.log("filter",a),n({page:h.value,pageSize:N.value,sortField:i.value,sortOrder:c.value,filters:a.value})};return pe(()=>{g()}),(e,a)=>{const v=s("a-page-header"),f=s("a-input"),F=s("a-form-item"),y=s("a-button"),X=s("a-form"),j=s("a-divider"),q=s("a-space"),Q=s("a-table"),W=s("a-modal");return A(),me("div",_e,[o(v,{style:{border:"1px solid rgb(235 237 240)"},"sub-title":"Role list page",title:"Roles"}),o(X,{ref_key:"searchForm",ref:E,model:t.value,layout:"inline"},{default:l(()=>[o(F,{label:"Name"},{default:l(()=>[o(f,{value:t.value.Name,"onUpdate:value":a[0]||(a[0]=u=>t.value.Name=u),placeholder:"Name"},null,8,["value"])]),_:1}),o(F,{label:"Description"},{default:l(()=>[o(f,{value:t.value.Description,"onUpdate:value":a[1]||(a[1]=u=>t.value.Description=u),placeholder:"Description"},null,8,["value"])]),_:1}),o(F,null,{default:l(()=>[o(y,{type:"primary",onClick:P},{icon:l(()=>[o(r(ne))]),default:l(()=>[a[3]||(a[3]=k(" Search "))]),_:1}),o(y,{icon:b(r(ue)),type:"primary",onClick:U},{default:l(()=>a[4]||(a[4]=[k(" Reset ")])),_:1},8,["icon"]),o(r(I),{type:"vertical"}),o(y,{icon:b(r(de)),type:"primary",onClick:G},{default:l(()=>a[5]||(a[5]=[k(" New ")])),_:1},8,["icon"])]),_:1})]),_:1},8,["model"]),o(j,{type:"horizontal"}),o(Q,{columns:O,"data-source":r(p)?r(p).list:null,loading:r(_),pagination:B.value,run:r(n),onChange:H},{bodyCell:l(({column:u,record:z})=>[u.key==="action"?(A(),ve(q,{key:0},{default:l(()=>[o(y,{icon:b(r(re)),type:"primary",onClick:Y=>V(z)},{default:l(()=>a[6]||(a[6]=[k(" Edit ")])),_:2},1032,["icon","onClick"]),o(r(I),{type:"vertical"}),o(y,{icon:b(r(ie)),danger:"",onClick:Y=>J(z.id,z.name)},{default:l(()=>a[7]||(a[7]=[k(" Delete ")])),_:2},1032,["icon","onClick"])]),_:2},1024)):fe("",!0)]),_:1},8,["data-source","loading","pagination","run"]),o(W,{open:m.value,"onUpdate:open":a[2]||(a[2]=u=>m.value=u),"mask-closable":!1,title:"Edit Role Info",width:"50%",onCancel:L,onOk:K},{default:l(()=>[o(le,{ref_key:"modalForm",ref:R,"form-model":S.value},null,8,["form-model"])]),_:1},8,["open"])])}}});export{Re as default};