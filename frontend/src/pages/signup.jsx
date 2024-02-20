export default function Signup(){
    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-80 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Fill out your INFYnaut Application!</h1>
            </div>
            <br/>
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <form>
                    <label for="username">Username:</label><br/>
                    <input type="text" id="username" name="username"></input><br/><br/>
                    <label for="password">Password:</label><br/>
                    <input type="text" id="password" name="password"></input><br/><br/>
                    <label for="email">Email:</label><br/>
                    <input type="text" id="email" name="email"></input><br/><br/>
                    <input type="submit" value="Submit" className="bg-violet-900" />
                </form>
            </div>
        </div>
    )
}