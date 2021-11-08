export default {
  theme: 'snow',
  boundary: document.body, 
  modules: {
    toolbar: [
      ['bold', 'italic', 'underline', 'strike'],
      ['blockquote', 'code-block'],
      [{ 'header': 1 }, { 'header': 2 }],
      [{ 'list': 'ordered' }, { 'list': 'bullet' }],
      [{ 'script': 'sub' }, { 'script': 'super' }],
      [{ 'indent': '-1' }, { 'indent': '+1' }],
      [{ 'direction': 'rtl' }],
      [{ 'size': ['10px', '12px', '14px', '16px' ,'18px', '20px', '22px', '24px', '26px', '32px', '48px',false] }],
      [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
      [{ 'color': [] }, { 'background': [] }],
      [{ 'font': [] }],
      [{ 'align': [] }],
      ['clean'],
      ['link', 'image'],
	  [{ 'lineheight': ["1", "1.5", "1.75", "2", "3", "4", "5"] }], // 对齐方式
    ]
  },
  placeholder: 'Insert text here ...',
  readOnly: false
}
