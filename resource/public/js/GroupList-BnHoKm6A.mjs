var x=(E,O,i)=>new Promise((c,o)=>{var C=n=>{try{p(i.next(n))}catch(_){o(_)}},h=n=>{try{p(i.throw(n))}catch(_){o(_)}},p=n=>n.done?c(n.value):Promise.resolve(n.value).then(C,h);p((i=i.apply(E,O)).next())});import{u as Z}from"./index.es-CT2W4LN0.mjs";import{g as ee,d as ae,s as te}from"./UserTable.vue_vue_type_script_setup_true_lang-DBrY9MZf.mjs";import{n as oe}from"./util-oQdYwX59.mjs";import{S as M}from"./enum-BLvisX2J.mjs";import{_ as le}from"./GroupForm.vue_vue_type_script_setup_true_lang-BQG2I4ks.mjs";import{S as ne,D as I,E as re,e as ie,M as se}from"./bootstrap-5eopxyZm.mjs";import{R as ue,M as de}from"./RollbackOutlined-C4ajMdEa.mjs";import{a6 as ce,P as d,J as $,a7 as pe,aI as s,aF as A,aJ as me,n as t,aH as l,aa as r,ao as k,h as b,aG as ve,aX as fe}from"../jse/index-index-CmiOfTxA.mjs";import"./user-7A8qe_c7.mjs";const _e={class:"p-5"},Fe=ce({__name:"GroupList",setup(E){const O=[{title:"Id",dataIndex:"id",key:"Id",sorter:!0,defaultSortOrder:"ascend"},{title:"Name",dataIndex:"name",key:"Name",name:"Name",sorter:!0,filtered:!0},{title:"Description",dataIndex:"description",key:"Description",filtered:!0},{title:"Actions",key:"action"}],i=d("id"),c=d(M.Asc),o=d({Name:"",Description:""}),C=$(()=>{const e=[];return o.value.Name!==""&&e.push("name like ?"),o.value.Description!==""&&e.push("description like ?"),e}),h=$(()=>{const e=[];return o.value.Name!==""&&e.push(`%${o.value.Name}%`),o.value.Description!==""&&e.push(`%${o.value.Description}%`),e}),{data:p,run:n,loading:_,total:R,current:D,pageSize:N}=Z(ee,{defaultParams:[{page:1,pageSize:10,sortField:"id",sortOrder:M.Asc,filters:C.value,args:h.value}],pagination:{currentKey:"page",pageSizeKey:"pageSize"}}),T=$(()=>({total:R.value,current:D.value,pageSize:N.value})),G=d(),g=()=>{setTimeout(()=>{n({page:D.value,pageSize:N.value,sortField:i.value,sortOrder:c.value,filters:C.value,args:h.value})},500)},B=()=>{console.log("searchModel",o.value),g()},P=()=>{o.value={Name:"",Description:""},G.value.resetFields(),g()},m=d(!1),S=d(),F=d(),U=e=>{S.value=e,m.value=!0},V=(e,a)=>{se.confirm({title:`Deleting Group ${e}`,content:`Are you sure delete this group: ${a}?`,onOk(){ae(e).then(()=>g())},onCancel(){console.log("Cancel")}}),console.log("record",e)},J=()=>x(this,null,function*(){try{const e=yield F.value.validate();if(!e)return;console.log("form validate",e)}catch(e){console.log("error submit:",e);return}te(S.value),m.value=!1,g()}),K=()=>{F.value.resetForm(),m.value=!1},L=()=>x(this,null,function*(){const e=yield oe();S.value={id:e.id,name:"",description:""},m.value=!0}),H=(e,a,v)=>{if(D.value=e.current,N.value=e.pageSize,v.field){const f=v.field;i.value=f!=null?f:v.field,c.value=v.order}else i.value="id",c.value=M.Asc;n({page:D.value,pageSize:N.value,sortField:i.value,sortOrder:c.value,filters:a.value})};return pe(()=>{g()}),(e,a)=>{const v=s("a-page-header"),f=s("a-input"),z=s("a-form-item"),y=s("a-button"),X=s("a-form"),j=s("a-divider"),q=s("a-space"),Q=s("a-table"),W=s("a-modal");return A(),me("div",_e,[t(v,{style:{border:"1px solid rgb(235 237 240)"},"sub-title":"Group list page",title:"Groups"}),t(X,{ref_key:"searchForm",ref:G,model:o.value,layout:"inline"},{default:l(()=>[t(z,{label:"Name"},{default:l(()=>[t(f,{value:o.value.Name,"onUpdate:value":a[0]||(a[0]=u=>o.value.Name=u),placeholder:"Name"},null,8,["value"])]),_:1}),t(z,{label:"Description"},{default:l(()=>[t(f,{value:o.value.Description,"onUpdate:value":a[1]||(a[1]=u=>o.value.Description=u),placeholder:"Description"},null,8,["value"])]),_:1}),t(z,null,{default:l(()=>[t(y,{type:"primary",onClick:B},{icon:l(()=>[t(r(ne))]),default:l(()=>[a[3]||(a[3]=k(" Search "))]),_:1}),t(y,{icon:b(r(ue)),type:"primary",onClick:P},{default:l(()=>a[4]||(a[4]=[k(" Reset ")])),_:1},8,["icon"]),t(r(I),{type:"vertical"}),t(y,{icon:b(r(de)),type:"primary",onClick:L},{default:l(()=>a[5]||(a[5]=[k(" New ")])),_:1},8,["icon"])]),_:1})]),_:1},8,["model"]),t(j,{type:"horizontal"}),t(Q,{columns:O,"data-source":r(p)?r(p).list:null,loading:r(_),pagination:T.value,run:r(n),onChange:H},{bodyCell:l(({column:u,record:w})=>[u.key==="action"?(A(),ve(q,{key:0},{default:l(()=>[t(y,{icon:b(r(re)),type:"primary",onClick:Y=>U(w)},{default:l(()=>a[6]||(a[6]=[k(" Edit ")])),_:2},1032,["icon","onClick"]),t(r(I),{type:"vertical"}),t(y,{icon:b(r(ie)),danger:"",onClick:Y=>V(w.id,w.name)},{default:l(()=>a[7]||(a[7]=[k(" Delete ")])),_:2},1032,["icon","onClick"])]),_:2},1024)):fe("",!0)]),_:1},8,["data-source","loading","pagination","run"]),t(W,{open:m.value,"onUpdate:open":a[2]||(a[2]=u=>m.value=u),"mask-closable":!1,title:"Edit Group Info",width:"50%",onCancel:K,onOk:J},{default:l(()=>[t(le,{ref_key:"modalForm",ref:F,"form-model":S.value},null,8,["form-model"])]),_:1},8,["open"])])}}});export{Fe as default};