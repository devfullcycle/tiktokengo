use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn get_qtd_tokens(prompt: *const libc::c_char) -> libc::c_uint {
    let prompt = unsafe { CStr::from_ptr(prompt).to_str().unwrap() };
    let count = tiktoken_rs::cl100k_base_singleton().lock().encode_with_special_tokens(prompt).len();
    count as libc::c_uint
}