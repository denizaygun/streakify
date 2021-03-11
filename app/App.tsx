import { StatusBar } from 'expo-status-bar';
import React, { useEffect, useState } from 'react';
import { SafeAreaView, View, ActivityIndicator, FlatList, StyleSheet, Text, TextInput } from 'react-native';

const host = "http://192.168.0.33:8080"

const Item = ({ name, icon, count }) => (
    <View style={styles.item}>
        <View style={{ flex: 0.1 }}>
            <Text>{icon}</Text>
        </View>
        <View style={{ flex: 0.7 }}>
            <Text style={styles.name}>{name}</Text>
        </View>
        <View style={{ flex: 0.2 }}>
            <Text>🔥 {count}</Text>
        </View>
    </View>

);

const App = () => {
    const [isLoading, setLoading] = useState(true);
    const [data, setData] = useState([]);

    const renderItem = ({ item }) => (
        <Item name={item.name} icon={item.icon} count={item.count} />
    );

    useEffect(() => {
        fetch(`${host}/streaks`)
            .then((response) => response.json())
            .then((json) => setData(json.streaks))
            .catch((error) => console.error(error))
            .finally(() => setLoading(false));
    }, []);

    return (
        <SafeAreaView style={styles.container}>
            <Text style={{ fontSize: 28, padding: 20 }}>Streakify</Text>

            <TextInput
                style={{ height: 80, padding: 20, fontSize: 18 }}
                placeholder="add new streak..."
            />
            {isLoading ? <ActivityIndicator /> : (
                <FlatList
                    data={data}
                    renderItem={renderItem}
                    keyExtractor={({ id }, index) => id}
                />
            )}
            <StatusBar style="auto" />
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        marginTop: StatusBar.currentHeight || 0,
    },
    item: {
        flexDirection: "row",
        backgroundColor: '#F8F8F8',
        padding: 20,
        marginVertical: 8,
        marginHorizontal: 16,
    },
    name: {
        fontSize: 18,
    },
});

export default App;