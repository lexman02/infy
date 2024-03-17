import {HeartIcon, ChatBubbleOvalLeftIcon} from "@heroicons/react/20/solid";
import defaultAvatar from "../img/default-avatar.png";

const fullName = (user) => {
    return `${user.profile.first_name} ${user.profile.last_name}`;
}

export default function Post({ post }) {
    return (
        <div className="flex justify-between bg-black/40 p-4 text-neutral-100">
            <div className="flex flex-col justify-around">
                {/* Post author details */}
                <div className="flex space-x-2 items-center">
                    <img src={`${post.user.profile.avatar || defaultAvatar}`} alt={fullName(post.user)} className="w-11 h-11 rounded-full" />
                    <div>
                        <div className="flex items-end space-x-1">
                            <h2 className="font-bold">
                                {fullName(post.user)}
                            </h2>
                            <span className="text-neutral-500 text-sm font-light">@{post.user.username}</span>
                        </div>
                        <div className="flex items-end space-x-1 text-sm text-neutral-400">
                            <p className="font-light">
                                watched
                            </p>
                            <span className="font-medium">{post.movie.title}</span>
                        </div>
                    </div>
                </div>
                {/* Post content */}
                <h2 className="font-medium">
                    {post.content}
                </h2>
                {/* Post interaction buttons */}
                <div className="flex space space-x-4">
                    <HeartIcon className="h-6 w-6 text-neutral-600" />
                    <ChatBubbleOvalLeftIcon className="h-6 w-6 text-neutral-200" />
                </div>
            </div>
            <img src={`https://image.tmdb.org/t/p/original/${post.movie.poster_path}`} alt={post.movie.title} className="w-20 h-32 object-cover rounded-lg" />
        </div>
    )
}