import About from "./About";
import Instgram from "./Instgram";

export default function SideBar() {
    return (
        <>
            <aside className="w-full  flex flex-col items-center px-3">

                <About />

                <Instgram name={'Instgram'} link={"https://ajf.sa/a"} />
               

            </aside>
        </>
    )
}