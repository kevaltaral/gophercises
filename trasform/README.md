# Exercise 18: Image Transformer


## Exercise details

Create a web server where the user uploads an image, and then is guided through a selection process using various options on the [primitive](https://github.com/fogleman/primitive) CLI.

For instance, a user might upload an image and then be present with outputs using several of the different modes:

```
mode: 0=combo, 1=triangle, 2=rect, 3=ellipse, 4=circle, 5=rotatedrect, 6=beziers, 7=rotatedellipse, 8=polygon
```

A user could then selec the image with the mode they prefer, at which time the web server would then assume the user prefers that mode and give them new options with that mode and another variable, such as the number of shapes (the `n` flag). At this time the server would output maybe 6 samples each using a different `n` value on the original image with the selected mode and the user could choose the they prefer.

While primitive doesn't have a ton of options, there are enough to at least create 2-3 steps like described above and then once all the settings are selected you could produce perhaps 4 images using those settings and let the user choose their favorite amongst those and download it.
