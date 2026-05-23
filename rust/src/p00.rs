pub fn run () {
   let a: &str = "パトカー";
   let b: &str="タクシー";

    for i in 0..a.len() as usize {
        print!("{}{}", a.chars().nth(i).unwrap(), b.chars().nth(i).unwrap());
    }
    print!("\n");
}