use std::env;

mod p00;

fn main() {
   let args: Vec<String> = env::args().collect();

   if args[1] == "p00" {
        p00::run();
   }
}
