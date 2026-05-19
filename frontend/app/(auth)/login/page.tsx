import FormLogin from "@/components/auth/FormLogin";
import Link from "next/link";

export default function LoginPage() {
  return (
    <div className="grid min-h-screen grid-cols-1 md:grid-cols-2">
      <div className="bg-primary hidden h-full w-full flex-col justify-between p-6 md:flex md:p-8 lg:p-12 xl:p-16">
        <div className="">
          <h1 className="text-3xl font-bold text-white">HRIS Connect</h1>
          <div className="mt-12 space-y-6">
            <h4 className="text-lg text-white">
              Empowering People, Scaling Excellence.
            </h4>
            <p className="text-base text-white">
              Our mission is to simplify complex human resource processes
              through intuitive technology, allowing organizations to focus on
              what matters most: their people.
            </p>
          </div>
        </div>
        <div className="rounded-xl border border-white/50 bg-white/20 p-6">
          <div className="flex items-stretch gap-3">
            <div className="w-10 rounded-full bg-gray-400"></div>
            <div className="">
              <p className="text-base text-white">Sarah Jenkins</p>
              <p className="text-sm text-gray-200/80">
                Chief People Officer, InnovateCorp
              </p>
            </div>
          </div>
          <p className="mt-4 text-white italic">
            &quot;HRIS Connect transformed how we manage our global workface.
            Efficiency has increased by 40% since implementation.&quot;
          </p>
        </div>
      </div>
      {/* <div className='h-full w-full bg-white p-6 md:py-8 lg:mx-auto lg:w-3/4 lg:py-16 xl:w-1/2'> */}
      <div className="flex h-full w-full flex-col justify-center gap-10 bg-white p-6 md:justify-between md:gap-0 md:py-8 lg:mx-auto lg:w-3/4 lg:p-12 xl:p-16">
        <div className="">
          <h2 className="text-center text-2xl font-semibold text-black md:text-start">
            Welcome Back
          </h2>
          <p className="text-center text-base text-black md:text-start">
            Enter your credential to access the Admin Portal
          </p>
        </div>
        <FormLogin />
        <div className="flex w-full items-center justify-between">
          <p className="text-secondary text-sm">&copy; 2026 HRIS Connect</p>
          <Link href={`#`}>
            <p className="text-secondary text-sm">Privacy Policy</p>
          </Link>
          <Link href={`#`}>
            <p className="text-secondary text-sm">Terms of Service</p>
          </Link>
        </div>
      </div>
    </div>
  );
}
