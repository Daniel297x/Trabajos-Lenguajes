/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package javaapplication9;

 import javax.swing.*;
import java.awt.*;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.util.ArrayList;
import java.util.Arrays;
import static java.util.Arrays.stream;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import javax.swing.border.Border;

public class Proyecto1Daniel extends JFrame {
     
    private final int FILAS = 20;
    private final int COLUMNAS = 24;
    int[][] matriz = new int[FILAS][COLUMNAS];
    private final Color COLOR_CASILLA = Color.BLACK;
    private final Color COLOR_SELECCIONADO = Color.BLUE;
    Border border = BorderFactory.createLineBorder(Color.GREEN, 1);

    private JPanel panelGrid;
    private int filaSeleccionada = -1;
    private int columnaSeleccionada = -1;
    List<List<Integer>> listaObstaculos = new ArrayList<>();
    List<List<Integer>> listaObjetivos = new ArrayList<>();
   boolean start = false;
    

    public Proyecto1Daniel() {
        setTitle("Grid Interfaz");
        setSize(400, 400);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setLocationRelativeTo(null);

        panelGrid = new JPanel(new GridLayout(FILAS, COLUMNAS));
        panelGrid.setPreferredSize(new Dimension(300, 300));
        inicializarCasillas();

        add(panelGrid);
        
        panelGrid.addMouseListener(new MouseAdapter() {
            @Override
            public void mouseClicked(MouseEvent e) 
            {
                super.mouseClicked(e);
                int fila = e.getY() / (panelGrid.getHeight() / FILAS);
                int columna = e.getX() / (panelGrid.getWidth() / COLUMNAS);
                if((fila == 19 && columna == 23))
                    start = true;
                else if (fila >= 0 && fila < FILAS-2 && columna >= 0 && columna < COLUMNAS-4) {
                    filaSeleccionada = fila;
                    columnaSeleccionada = columna;
                    System.out.println( fila + "," + columna);
                    if(start)
                    {
                        if(!listaObstaculos.contains(Arrays.asList(fila, columna)))
                            listaObjetivos.add(Arrays.asList(fila, columna));  
                        actualizarCasillas();
                    }  
                    else if(!start)
                    {
                        listaObstaculos.add(Arrays.asList(fila, columna));
                        actualizarCasillas();
                    }
                } 
            }
        });
      
       
    }

    private void inicializarCasillas() 
    {
        int filas = matriz.length;
        int columnas = matriz[0].length;
        for (int i = 0; i < filas; i++) {
            for (int j = 0; j < columnas; j++) {
                JPanel casilla = new JPanel();
                if(i < 18 && j < 20)
                {
                casilla.setBackground(COLOR_CASILLA);
                casilla.setBorder(border);
                }
                else
                {
                    casilla.setBackground(Color.GRAY);
                }
                if(i == 19 && j == 23)
                    casilla.setBackground(Color.RED);
                panelGrid.add(casilla);
            }
        }

    }
    
    private void actualizarCasillas() {
        Component[] casillas = panelGrid.getComponents();
        int filas = matriz.length;
        int columnas = matriz[0].length;
        int total = 0;
        for (int i = 0; i < filas; i++) {
            for (int j = 0; j < columnas; j++) {
                 JPanel casilla = (JPanel) casillas[total];
                 total += 1;
                if((i < 18 && j < 20))
                {
                    casilla.setBackground(COLOR_CASILLA);
                    casilla.setBorder(border);
                }
                else
                    casilla.setBackground(Color.GRAY);
                for(int p = 0; p < listaObjetivos.size(); p++)
                {
                    if(i == listaObjetivos.get(p).get(0) && j == listaObjetivos.get(p).get(1))
                    casilla.setBackground(Color.CYAN); 
                }
                for(int p = 0; p < listaObstaculos.size(); p++)
                {
                    if(i == listaObstaculos.get(p).get(0) && j == listaObstaculos.get(p).get(1))
                    casilla.setBackground(COLOR_SELECCIONADO); 
                }
                
                if(i == 19 && j == 23)
                    casilla.setBackground(Color.RED);
            }
        }
    }
    
    public static List<List<Integer>> Vecinos(List<Integer> pos)
    {
        return Arrays.asList(arriba(pos),abajo(pos), izquierda(pos), derecha(pos));
    }
    
    public static List<Integer> arriba(List<Integer> pos)
    {
        
        if(pos.get(0) > 0 )
        {
            System.out.println(Arrays.asList(pos.get(0)-1,pos.get(1)));
            return Arrays.asList(pos.get(0)-1,pos.get(1));
        }
        else
            return pos;
    }
    
    public static List<Integer> abajo(List<Integer> pos)
    {
        if(pos.get(0) < 17 )
            return Arrays.asList(pos.get(0)+1,pos.get(1));
        else
            return pos;
    }
    
    public static List<Integer> derecha(List<Integer> pos)
    {
        if(pos.get(1) < 19 )
            return Arrays.asList(pos.get(0),pos.get(1)+1);
        else
            return pos;
    }
    
    public static List<Integer> izquierda(List<Integer> pos)
    {
        if(pos.get(0) > 0 )
            return Arrays.asList(pos.get(0),pos.get(1)-1);
        else
            return pos;
    }
    
    public static void GenRutaProfAux(List<Integer> pos, List<Integer> objetivo)
    {
        GenRutaProf(Arrays.asList(Arrays.asList(pos)), objetivo);
        
    }
    
    public static List<List<List<Integer>>>  GenRutaProf(List<List<List<Integer>>> lista, List<Integer> objetivo)
    {
        if(lista.isEmpty())
            return Arrays.asList();
        else if(solucion(lista.get(0).get(0), objetivo))
        {
            //return GenRutaProf((ArrayList)(Arrays.asList( ((ArrayList)lista.get(0).stream().sorted(Collections.reverseOrder()).collect(Collectors.toList())),((ArrayList)(extender(lista.get(0)).stream().flatMap(List::stream).collect(Collectors.toList()))),((ArrayList)(lista.subList(1,lista.size()).stream().flatMap(List::stream).collect(Collectors.toList())))).stream().filter(x -> !x.isEmpty()).collect(Collectors.toList())), objetivo);
             //arreglar  
                return GenRutaProf(Arrays.asList(((ArrayList)
                        (extender(lista.get(0)).stream()
                                .flatMap(List::stream)
                                .collect(Collectors.toList()))),
                        ((ArrayList)(lista.subList(1,lista.size())
                                .stream().flatMap(List::stream)
                                .collect(Collectors.toList())))), objetivo);

        }
        else
             return GenRutaProf(Arrays.asList(((ArrayList)(extender(lista.get(0))
                     .stream().flatMap(List::stream)
                     .collect(Collectors.toList()))),((ArrayList)
                             (lista.subList(1,lista.size())
                                     .stream().flatMap(List::stream)
                                     .collect(Collectors.toList())))), objetivo);
           
}
    
    public static boolean solucion(List<Integer> pos, List<Integer> objetivo)
    {
        if(pos.equals(objetivo))
            return true;
        else
            return false;
    }
    
    public static List<List<List<Integer>>> extender(List<List<Integer>> ruta)
    {
        
        return remIF((ArrayList)(Vecinos(ruta.get(0)).stream().map(x -> 
        {if(ruta.contains(x ) == true){ return Arrays.asList() ;}
        else return Arrays.asList((ArrayList)(Stream.concat(Stream.of(x),Stream.of(ruta))
                .flatMap(List::stream).collect(Collectors.toList()))) ;})
                .collect(Collectors.toList())));
         
    }
    
    public static List<List<List<Integer>>> remIF(List<List<Integer>> lista)
    {
        return Arrays.asList((ArrayList)(lista.stream()
                .map(x -> {if(x.isEmpty()) 
                    return Arrays.asList(); else return Arrays.asList(x);})
                .flatMap(List::stream).collect(Collectors.toList())));
          
    }
    public static void main(String[] args) {
        SwingUtilities.invokeLater(() -> {
            Proyecto1Daniel gridInterfaz = new Proyecto1Daniel();
            gridInterfaz.setVisible(true);
        });
        //System.out.println(extender(Arrays.asList(Arrays.asList(5,6))));
        GenRutaProfAux(Arrays.asList(5, 6), Arrays.asList(1, 3));
        
    }
}
