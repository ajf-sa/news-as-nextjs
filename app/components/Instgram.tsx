export default function Instgram({name,link}) {

    return (
        <div className="w-full bg-white shadow flex flex-col my-4 p-6">
            <p className="text-xl font-semibold pb-5">{name}</p>
            <div className="grid grid-cols-3 gap-3">
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=1" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=2" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=3" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=4" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=5" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=6" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=7" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=8" />
                <img className="hover:opacity-75" src="https://source.unsplash.com/collection/1346951/150x150?sig=9" />
            </div>
            <a href={link} className="w-full bg-cyan-700 text-white font-bold text-sm uppercase rounded hover:bg-cyan-800 flex items-center justify-center px-2 py-3 mt-6" style={{ direction: "ltr" }}>
                تابعنا @ajf.sa
            </a>
        </div>
    )
}