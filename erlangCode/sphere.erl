-module(sphere).
-compile([export_all]).

volume(Radius, AllResults_PID) ->
    io:format("Volume Calculation begins~n",[]),
    Volume = (4/3) * math:pi() * Radius * Radius * Radius,
    AllResults_PID ! {volume, self()},
    receive
        okay ->
            io:format("Volume is ~p ~n",[Volume])
    end.

surface(Radius, AllResults_PID) ->
    io:format("Surface Calculation begins~n",[]),
    Surface = 4.0 * math:pi() * math:sqrt(Radius),
    AllResults_PID ! {surface, self()},
    receive
        okay ->
            io:format("Surface is ~p ~n",[Surface])
    end.

circularArea(Radius, AllResults_PID) ->
    io:format("Circular Area Calculation begins~n",[]),
    Area = math:pi() * math:sqrt(Radius),
    AllResults_PID ! {circularArea, self()},
    receive
        okay ->
            io:format("Circular Area is ~p ~n",[Area])
    end.

allResults() ->
    receive
        {volume, Volume_PID} ->
            io:format("Volume was calculated~n"),
            Volume_PID ! okay,
            allResults();
        {surface, Surface_PID} ->
            io:format("Surface was calculated~n"),
            Surface_PID ! okay,
            allResults();
        {circularArea, Area_PID} ->
            io:format("Circular Area was calculated~n"),
            Area_PID ! okay,
            allResults()
    end.

start(Radius) ->
    AllResults_PID=spawn(sphere, allResults, []),
    spawn(sphere, volume, [Radius, AllResults_PID]),
    spawn(sphere, circularArea, [Radius, AllResults_PID]),
    spawn(sphere, surface, [Radius, AllResults_PID]).